package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	policyController "github.com/phamdinhha/go-authorizer/internal/policy/controller"
	policyService "github.com/phamdinhha/go-authorizer/internal/policy/service"
	roleController "github.com/phamdinhha/go-authorizer/internal/role/controller"
	roleService "github.com/phamdinhha/go-authorizer/internal/role/service"
)

func (s *Server) MapHandlers(app *fiber.App) error {
	e, err := casbin.NewEnforcer(s.cfg.Casbin.ModelConfig, s.cfg.Casbin.Policy)
	if err != nil {
		return err
	}

	policyService := policyService.NewPolicyService(s.cfg, s.logger, e)
	roleService := roleService.NewRoleService(s.cfg, s.logger, e)

	policyCtl := policyController.NewPolicyController(s.logger, s.cfg, policyService)
	roleCtl := roleController.NewRoleController(s.logger, s.cfg, roleService)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	health := app.Group("/health")
	policyGroup := app.Group("/policy")
	roleGroup := app.Group("/role")

	policyController.MapPolicyRoutes(policyGroup, policyCtl)
	roleController.MapRoleRoutes(roleGroup, roleCtl)

	health.Get("/", func(c *fiber.Ctx) error {
		s.logger.Infof("Health checking...")
		return c.JSON(fiber.StatusOK)
	})

	return nil
}
