package role

import (
	"context"

	"github.com/phamdinhha/go-authorizer/internal/models"
)

type RoleService interface {
	AddUserToRole(ctx context.Context, req *models.Role) (*models.Role, error)
	RemoveUserFromRole(ctx context.Context, req *models.Role) (*models.Role, error)
}
