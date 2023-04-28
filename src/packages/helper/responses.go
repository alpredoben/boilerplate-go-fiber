package helpers

import "github.com/gofiber/fiber/v2"

func GenerateResponse(ctx *fiber.Ctx, message string, statusCode int, data interface{}) error {
	if statusCode >= 100 || statusCode <= 308 {
		return ctx.Status(statusCode).JSON(fiber.Map{
			"success": true,
			"message": message,
			"data":    data,
		})
	}
	return ctx.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error":   data,
	})
}
