package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/authz"
	"github.com/phamdinhha/go-authorizer/internal/models"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
	"github.com/phamdinhha/go-authorizer/pkg/utils"
)

type authzController struct {
	logger       logger.Logger
	cfg          *config.Config
	authzService authz.AuthzService
}

func NewAuthzController(
	logger logger.Logger,
	cfg *config.Config,
	authzService authz.AuthzService,
) authz.AuthzController {
	return &authzController{logger: logger, cfg: cfg, authzService: authzService}
}

func (ctl *authzController) CheckAuthorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := models.Request{}
		if err := c.BodyParser(&req); err != nil {
			ctl.logger.Errorf("AuthzController error when parsing request body: %v", err)
			return utils.FiberReponse(c, fiber.StatusBadRequest, []interface{}{}, err)
		}
		ok, err := ctl.authzService.CheckAuthorization(c.Context(), &req)
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if !ok {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{"user does not authorized to access this resource"}, errors.New("unauthorized"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, "")
	}
}
