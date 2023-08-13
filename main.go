package main

import (
	"log"
  "context"
	"github.com/gofiber/fiber/v2"
	redis "github.com/redis/go-redis/v9"
)

func main(){
  ctx := context.Background()
  options, err := redis.ParseURL("redis://localhost:6379")
  if err!=nil {
    log.Fatal(err)
  }
  redis_client := redis.NewClient(options)
  app := fiber.New()

  app.Get("/info", func(c *fiber.Ctx) error {
    version := redis_client.Info(ctx)
    return c.SendString(version.String())
  })

  app.Listen(":8080")
}
