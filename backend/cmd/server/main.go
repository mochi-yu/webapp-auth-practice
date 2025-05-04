package main

import (
	"log"

	"github.com/mochi-yu/webapp-auth-practice/config"
	"github.com/mochi-yu/webapp-auth-practice/infrastructure"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	db, err := infrastructure.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := infrastructure.NewRouter(db, cfg)
	router.Run(":8080")
}
