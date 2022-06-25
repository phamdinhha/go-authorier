package policy

import "github.com/gofiber/fiber/v2"

type PolicyController interface {
	CreatePolicy() fiber.Handler
	UpdatePolicy() fiber.Handler
	DeletePolicy() fiber.Handler
	GetAllPolicies() fiber.Handler
}
