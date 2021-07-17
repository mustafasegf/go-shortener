package main

import (
	"log"

	"github.com/mustafasegf/go-shortener/util"
	"github.com/mustafasegf/go-shortener/api"
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

	server := api.MakeServer(config, db)
	server.RunServer()
}
