package main

import (
	"log"

	"gitlab.mapcard.pro/external-map-team/shop-service/config"
	"gitlab.mapcard.pro/external-map-team/shop-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
