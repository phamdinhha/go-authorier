package service

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/authz"
	"github.com/phamdinhha/go-authorizer/internal/models"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
)

type authzService struct {
	cfg      *config.Config
	logger   logger.Logger
	enforcer *casbin.Enforcer
}

func NewAuthzService(
	cfg *config.Config,
	logger logger.Logger,
	enforcer *casbin.Enforcer,
) authz.AuthzService {
	return &authzService{cfg: cfg, logger: logger, enforcer: enforcer}
}

func (s *authzService) CheckAuthorization(ctx context.Context, req *models.Request) (ok bool, err error) {
	ok, err = s.enforcer.Enforce(req.User, req.Object, req.Action)
	if err != nil {
		return ok, err
	}
	return ok, nil
}
