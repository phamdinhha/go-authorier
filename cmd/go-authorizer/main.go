package main

import (
	"log"
	"os"

	"github.com/phamdinhha/go-authorizer/config"
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

}
