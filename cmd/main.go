package main

import (
	"log"
	"transaction-service/config"
	"transaction-service/internal/app"
)



func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error config: %s", err)
	}

	log.Println("app success running")

	app.Run(cfg)
}