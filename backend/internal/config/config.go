package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
		POSTGRES    PostgresConfig
		NATS        NatsConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	PostgresConfig struct {
		Db       string `mapstructure:"db"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Dsn      string `mapstructure:"dsn"`
	}

	NatsConfig struct {
		Url       string `mapstructure:"url"`
		ClusterID string `mapstructure:"clusterID"`
		ClientID  string `mapstructure:"clientID"`
		Subject   string `mapstructure:"subject"`
	}
)

func Init(configsDir string) (*Config, error) {
	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.POSTGRES); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("nats", &cfg.NATS); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	if os.Getenv("DB_HOST") != "" {
		cfg.POSTGRES.Dsn = fmt.Sprintf("host=%s port=%s dbname=wbdb user=wbdb password=dbwb sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	}
	if os.Getenv("NATS_HOST") != "" {
		cfg.NATS.Url = fmt.Sprintf("nats://%s:%s", os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"))
	}
}
