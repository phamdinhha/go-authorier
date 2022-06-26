package authz

import (
	"context"

	"github.com/phamdinhha/go-authorizer/internal/models"
)

type AuthzService interface {
	CheckAuthorization(ctx context.Context, req *models.Request) (bool, error)
}
