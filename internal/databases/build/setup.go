package build

import (
	"github.com/go-redis/redis/v7"
	"github.com/kasfulk/golang-library/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	var err error

	var DBUser = config.DatabaseConfig().Username
	var DBPass = config.DatabaseConfig().Password
	var DBHost = config.DatabaseConfig().Host
	var DBPort = config.DatabaseConfig().Port
	var DBName = config.DatabaseConfig().Database

	ConnectionString := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8"
	db, err := gorm.Open(mysql.Open(ConnectionString), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}

func ConnectRedis() *redis.Client {
	var RedisURL = config.RedisConfig().Url
	var RedisPassword = config.RedisConfig().Password

	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisURL,
		Password: RedisPassword,
		DB:       0,
	})
	return rdb
}
