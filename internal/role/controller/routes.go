package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-authorizer/internal/role"
)

func MapRoleRoutes(roleGroup fiber.Group, ctl role.RoleController) {
	roleGroup.Post("/", ctl.AddUserToRole())
	roleGroup.Delete("/", ctl.RemoveUserFromRole())
}
