package service

import (
	"context"
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/models"
	"github.com/phamdinhha/go-authorizer/internal/policy"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
)

type policyService struct {
	cfg      *config.Config
	logger   logger.Logger
	enforcer *casbin.Enforcer
}

func NewPolicyService(
	cfg *config.Config,
	logger logger.Logger,
	enforcer *casbin.Enforcer,
) policy.PolicyService {
	return &policyService{cfg: cfg, logger: logger, enforcer: enforcer}
}

func (s *policyService) CreatePolicy(ctx context.Context, req *models.Policy) (*models.Policy, error) {
	isAdded, err := s.enforcer.AddPolicy(req.Subject, req.Object, req.Action)
	if err != nil {
		s.logger.Errorf("PolicyService: create policy error: %v", err)
		return nil, err
	}
	if !isAdded {
		return nil, nil
	}
	return req, nil
}

func (s *policyService) UpdatePolicy(ctx context.Context, oldPolicy *models.Policy, newPolicy *models.Policy) (*models.Policy, error) {
	oldP := []string{oldPolicy.Subject, oldPolicy.Object, oldPolicy.Action}
	newP := []string{newPolicy.Subject, newPolicy.Object, newPolicy.Action}
	isUpdated, err := s.enforcer.UpdatePolicy(oldP, newP)
	if err != nil {
		s.logger.Errorf("PolicyService: update policy error: %v", err)
		return nil, err
	}
	if !isUpdated {
		return nil, nil
	}
	return newPolicy, nil
}

func (s *policyService) DeletePolicy(ctx context.Context, req *models.Policy) (*models.Policy, error) {
	isDeleted, err := s.enforcer.RemovePolicy(req.Subject, req.Object, req.Action)
	if err != nil {
		s.logger.Errorf("PolicyService: delete policy error: %v", err)
		return nil, err
	}
	if !isDeleted {
		return nil, nil
	}
	return req, nil
}

func (s *policyService) GetAllPolicies(ctx context.Context) (res []models.Policy, err error) {
	policyList := s.enforcer.GetPolicy()
	for _, policy := range policyList {
		if len(policy) != 3 {
			s.logger.Errorf("PolicyService: wrong policy format found: %v", policy)
			return res, errors.New("wrong policy format found")
		}
		res = append(res, models.Policy{
			Subject: policy[0],
			Object:  policy[1],
			Action:  policy[2],
		})
	}
	return res, nil
}
