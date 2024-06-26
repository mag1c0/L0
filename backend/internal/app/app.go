package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cache2 "github.com/mag1c0/L0/backend/internal/cache"
	"github.com/mag1c0/L0/backend/internal/config"
	"github.com/mag1c0/L0/backend/internal/delivery/amqp"
	delivery "github.com/mag1c0/L0/backend/internal/delivery/http"
	"github.com/mag1c0/L0/backend/internal/generator"
	"github.com/mag1c0/L0/backend/internal/repository"
	"github.com/mag1c0/L0/backend/internal/server"
	"github.com/mag1c0/L0/backend/internal/service"
	"github.com/mag1c0/L0/backend/pkg/db/pg"
	"github.com/mag1c0/L0/backend/pkg/nats"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Run(ctx context.Context, configPath string) {
	// Init configs
	cfg, err := config.Init(configPath)
	if err != nil {
		fmt.Println("Error reading config:", err)
	}

	// Init clients
	pgClient, err := pg.New(ctx, cfg.POSTGRES.Dsn)
	if err != nil {
		log.Fatalf("Failed to create db client: %v", err)
	}
	err = pgClient.DB().Ping(ctx)
	if err != nil {
		log.Fatalf("Ping pgdb error: %s", err.Error())
	}
	fmt.Println("Postgresql connection success")

	nsClient, err := nats.New(cfg.NATS.Url, cfg.NATS.ClusterID, cfg.NATS.ClientID)
	if err != nil {
		log.Fatalf("Failed to create nats client: %v", err)
	}
	defer nsClient.Sc.Close()
	fmt.Println("Nats connection success")

	// Services, Repos & Handlers
	cache := cache2.New()
	repos := repository.NewRepositories(pgClient)
	services := service.NewServices(service.Deps{
		Repos: repos,
		Cache: cache,
	})
	handlers := delivery.NewHandler(services)
	consumers := amqp.NewConsumer(amqp.Deps{
		Services: services,
		Stan:     nsClient.Sc,
	})

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		if err := consumers.Orders.Subscribe(ctx, cfg.NATS.Subject); err != nil {
			fmt.Printf("Failed to subscribe to subject %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		for {
			order := generator.GenerateOrder()
			fmt.Println("Generated order:", order.OrderUID)
			orderJson, err := json.Marshal(order)
			if err != nil {
				fmt.Printf("Error marshalling order %v", err)
			}

			err = nsClient.Sc.Publish(cfg.NATS.Subject, orderJson)
			if err != nil {
				fmt.Printf("Error publish: %v\n", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init())
	go func() {
		defer wg.Done()

		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error occurred while running http server: %s\n", err.Error())
		}
	}()
	fmt.Println("HTTP Server started")

	wg.Wait()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(ctx, timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		fmt.Printf("Failed to stop server: %v", err)
	}
}
