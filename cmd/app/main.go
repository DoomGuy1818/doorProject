package main

import (
	"doorProject/internal/config/configs"
	"doorProject/internal/db"
	"doorProject/internal/server"
	"log"
)

func main() {
	client, err := db.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}

	var config configs.DatabaseConfig
	config.SetupDb(client)

	go server.Start()
	select {}
}
