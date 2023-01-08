package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to read config file")
	}
}

func ServerPort() string {
	cfg := viper.GetString("server.port")
	if cfg == "" {
		return ":8080"
	}

	return fmt.Sprintf(":%s", cfg)
}

func DatabaseConfig() DatabaseEnvironment {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	database := viper.GetString("database.dbname")

	return DatabaseEnvironment{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func RedisConfig() RedisEnvirontment {
	host := viper.GetString("redis.url")
	password := viper.GetString("redis.password")

	return RedisEnvirontment{
		Url:      host,
		Password: password,
	}
}
