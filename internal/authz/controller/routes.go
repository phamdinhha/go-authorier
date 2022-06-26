package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-authorizer/internal/authz"
)

func MapAuthzRoutes(authzGroup fiber.Router, ctl authz.AuthzController) {
	authzGroup.Post("/", ctl.CheckAuthorization())
}
