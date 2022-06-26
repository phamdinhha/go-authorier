package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-authorizer/internal/policy"
)

func MapPolicyRoutes(policyGroup fiber.Router, ctl policy.PolicyController) {
	policyGroup.Get("/all", ctl.GetAllPolicies())
	policyGroup.Post("/", ctl.CreatePolicy())
	policyGroup.Put("/", ctl.UpdatePolicy())
	policyGroup.Delete("/", ctl.DeletePolicy())
}
