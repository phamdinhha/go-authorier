package role

import "github.com/gofiber/fiber/v2"

type RoleController interface {
	AddUserToRole() fiber.Handler
	RemoveUserFromRole() fiber.Handler
}
