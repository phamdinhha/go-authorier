package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
)

type Server struct {
	fiber  *fiber.App
	cfg    *config.Config
	logger logger.Logger
	db     *sqlx.DB
}
