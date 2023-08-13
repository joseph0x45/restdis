package commands

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Result struct {
	data string
}

func Info(r *redis.Client, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		version, err := r.Info(ctx).Result()
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(fiber.Map{
			"data": version,
		})
	}
}

func Get(r *redis.Client, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")
		value, err := r.Get(ctx, key).Result()
		if err == redis.Nil {
			return c.JSON(fiber.Map{
				"data": nil,
			})
		}
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(fiber.Map{
			"data": value,
		})
	}
}
