package main

import (
	"log"
	"os"

	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/pkg/db/postgres"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
)

func main() {
	log.Println("Starting go-authorizer service")
	configPath := os.Getenv("config")
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("LoadConfig error: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal("ParseConfig: %v", err)
	}
	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to init postgres db: %s", err)
	} else {
		logger.Infof("Successfully connected to postgres, status: %v", db.Stats())
	}
	defer db.Close()
}
