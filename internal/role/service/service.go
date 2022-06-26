package service

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/models"
	"github.com/phamdinhha/go-authorizer/internal/role"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
)

type userRoleService struct {
	cfg      *config.Config
	logger   logger.Logger
	enforcer *casbin.Enforcer
}

func NewRoleService(
	cfg *config.Config,
	logger logger.Logger,
	enforcer *casbin.Enforcer,
) role.RoleService {
	return &userRoleService{cfg: cfg, logger: logger, enforcer: enforcer}
}

func (s *userRoleService) AddUserToRole(ctx context.Context, req *models.Role) (*models.Role, error) {
	isAdded, err := s.enforcer.AddGroupingPolicy(req.User, req.Group)
	if err != nil {
		s.logger.Errorf("UserRoleService failed to add user role: %v", err)
		return nil, err
	}
	if !isAdded {
		return nil, nil
	}
	return req, nil
}

func (s *userRoleService) RemoveUserFromRole(ctx context.Context, req *models.Role) (*models.Role, error) {
	isRemoved, err := s.enforcer.RemoveGroupingPolicy(req.User, req.Group)
	if err != nil {
		s.logger.Errorf("UserRoleService failed to add user role: %v", err)
		return nil, err
	}
	if !isRemoved {
		return nil, nil
	}
	return req, nil
}
