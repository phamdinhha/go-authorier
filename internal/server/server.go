package server

import (
	"fmt"
	"os"
	"os/signal"

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

func NewServer(
	cfg *config.Config,
	logger logger.Logger,
	db *sqlx.DB,
) *Server {
	return &Server{fiber: fiber.New(), cfg: cfg, logger: logger, db: db}
}

func (s *Server) Run() error {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.fiber.Shutdown(); err != nil {
			s.logger.Errorf("Server is not running! Error: %s", err)
		}
		close(idleConnsClosed)
	}()

	s.logger.Info("Mapping handlers...")
	if err := s.MapHandlers(s.fiber); err != nil {
		s.logger.Errorf("Error when mapping handlers: %s", err)
		return err
	}

	server := fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)

	if err := s.fiber.Listen(server); err != nil {
		s.logger.Errorf("Server is not running! Error: %s", err)
		return err
	}

	<-idleConnsClosed
	return s.fiber.Server().Shutdown()
}
