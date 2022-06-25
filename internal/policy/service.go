package policy

import (
	"context"

	"github.com/phamdinhha/go-authorizer/internal/models"
)

type PolicyService interface {
	CreatePolicy(ctx context.Context, req *models.Policy) (*models.Policy, error)
	UpdatePolicy(ctx context.Context, oldPolicy *models.Policy, newPolicy *models.Policy) (*models.Policy, error)
	DeletePolicy(ctx context.Context, req *models.Policy) (*models.Policy, error)
	GetAllPolicies(ctx context.Context) ([]models.Policy, error)
}
