package main

import (
	"flag"
	"log"

	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/server"
	"github.com/phamdinhha/go-authorizer/pkg/db/postgres"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
)

var (
	configPath = flag.String("config", "config/config", "Config file")
)

func main() {
	log.Println("Starting go-authorizer service")
	cfgFile, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("LoadConfig error: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
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
	s := server.NewServer(cfg, logger, db)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
