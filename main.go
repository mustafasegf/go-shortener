package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/mustafasegf/go-shortener/api"
	"github.com/mustafasegf/go-shortener/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := util.SetLogger()
	if err != nil {
		log.Fatal("cannot set logger: ", err)
	}

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.DBDsn,
	}))
	if err != nil {
		log.Fatal("canot load db: ", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})

	server := api.MakeServer(config, db, rdb)
	server.RunServer()
}
