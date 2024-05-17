package config

import (
	"github.com/spf13/viper"
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
