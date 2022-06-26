package authz

import "github.com/gofiber/fiber/v2"

type AuthzController interface {
	CheckAuthorization() fiber.Handler
}
