package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-authorizer/config"
	"github.com/phamdinhha/go-authorizer/internal/models"
	"github.com/phamdinhha/go-authorizer/internal/policy"
	"github.com/phamdinhha/go-authorizer/pkg/logger"
	"github.com/phamdinhha/go-authorizer/pkg/utils"
)

type policyController struct {
	logger        logger.Logger
	cfg           *config.Config
	policyService policy.PolicyService
}

func NewPolicyController(
	logger logger.Logger,
	cfg *config.Config,
	policyService policy.PolicyService,
) policy.PolicyController {
	return &policyController{logger: logger, cfg: cfg, policyService: policyService}
}

func (ctl *policyController) CreatePolicy() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := models.Policy{}
		if err := c.BodyParser(&req); err != nil {
			ctl.logger.Errorf("PolicyController error when parsing request body: %v", err)
			return utils.FiberReponse(c, fiber.StatusBadRequest, []interface{}{}, err)
		}
		res, err := ctl.policyService.CreatePolicy(c.Context(), &req)
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if res == nil {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, errors.New("policy already created"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{res}, "")
	}
}

func (ctl *policyController) UpdatePolicy() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := models.UpdatePolicyReq{}
		if err := c.BodyParser(&req); err != nil {
			ctl.logger.Errorf("PolicyController error when parsing request body: %v", err)
			return utils.FiberReponse(c, fiber.StatusBadRequest, []interface{}{}, err)
		}
		res, err := ctl.policyService.UpdatePolicy(c.Context(), &req.OldPolicy, &req.NewPolicy)
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if res == nil {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, errors.New("policy cannot be updated"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{res}, "")
	}
}

func (ctl *policyController) DeletePolicy() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := models.Policy{}
		if err := c.BodyParser(&req); err != nil {
			ctl.logger.Errorf("PolicyController error when parsing request body: %v", err)
			return utils.FiberReponse(c, fiber.StatusBadRequest, []interface{}{}, err)
		}
		res, err := ctl.policyService.DeletePolicy(c.Context(), &req)
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if res == nil {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, errors.New("policy already deleted"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{res}, "")
	}
}

func (ctl *policyController) GetAllPolicies() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := ctl.policyService.GetAllPolicies(c.Context())
		if err != nil {
			return utils.FiberReponse(c, fiber.StatusInternalServerError, []interface{}{}, err)
		}
		if len(res) == 0 {
			return utils.FiberReponse(c, fiber.StatusOK, []interface{}{}, errors.New("empty policy list"))
		}
		return utils.FiberReponse(c, fiber.StatusOK, []interface{}{res}, "")
	}
}
