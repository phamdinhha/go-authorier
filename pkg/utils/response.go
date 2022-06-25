package utils

import "github.com/gofiber/fiber/v2"

func FiberReponse(c *fiber.Ctx, code int, data []interface{}, err interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"code":  code,
		"data":  data,
		"error": err,
	})
}
