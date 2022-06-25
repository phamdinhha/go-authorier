package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/models"
	"github.com/phamdinhha/go-authorizer/internal/role"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
	"github.com/phamdinhha/go-authorizer/pkg/utils"
)

type roleController struct {
	logger      logger.Logger
	cfg         *config.Config
	roleService role.RoleService
}

func NewRoleController(
	logger logger.Logger,
	cfg *config.Config,
	roleService role.RoleService,
) role.RoleController {
	return &roleController{logger: logger, cfg: cfg, roleService: roleService}
}

func (ctl *roleController) AddUserToRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := models.Role{}
		if err := c.BodyParser(&req); err != nil {
			ctl.logger.Errorf("RoleController error when parsing request body: %v", err)
			return utils.FiberReponse(c, fiber.StatusBadRequest, []interface{}{}, err)
		}
		res, err := ctl.roleService.AddUserToRole(c.Context(), &req)
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if res == nil {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, errors.New("role already added"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{res}, "")
	}
}

func (ctl *roleController) RemoveUserFromRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := models.Role{}
		if err := c.BodyParser(&req); err != nil {
			ctl.logger.Errorf("RoleController error when parsing request body: %v", err)
			return utils.FiberReponse(c, fiber.StatusBadRequest, []interface{}{}, err)
		}
		res, err := ctl.roleService.RemoveUserFromRole(c.Context(), &req)
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if res == nil {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, errors.New("role already added"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{res}, "")
	}
}
