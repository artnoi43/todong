package config

import (
	"fmt"
	"log"

	"github.com/go-yaml/yaml" // For logging
	"github.com/spf13/viper"

	"github.com/artnoi43/todong/domain/usecase/middleware"
	"github.com/artnoi43/todong/domain/usecase/redisclient"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/postgres"
)

type Config struct {
	Address    string             `mapstructure:"listen_address" yaml:"port"`
	ServerType enums.ServerType   `mapstructure:"server" yaml:"server"`
	Store      enums.StoreType    `mapstructure:"store" yaml:"store"`
	Middleware middleware.Config  `mapstructure:"middleware" yaml:"middleware"`
	Postgres   postgres.Config    `mapstructure:"postgres" yaml:"postgres"`
	Redis      redisclient.Config `mapstructure:"redis" yaml:"redis"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetDefault("listen_address", "127.0.0.1:8000")
	viper.SetDefault("store", "REDIS")
	viper.SetDefault("middleware.secret_key", "secret")
	viper.SetDefault("postgres.host", "127.0.0.1")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.password", "postgres")
	// viper.SetDefault("postgres.name", "postgres")
	viper.SetDefault("redis.db", 5)

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	if !config.Store.IsValid() {
		return nil, fmt.Errorf("invalid store type: %s", config.Store)
	}
	conf, _ := yaml.Marshal(config)
	log.Printf("\ntodogin Configuration:\n%s\n", conf)
	return config, nil
}
