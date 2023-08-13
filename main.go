package main

import (
	"log"
  "context"
	"github.com/gofiber/fiber/v2"
	redis "github.com/redis/go-redis/v9"
  "thewisepigeon/restdis/commands"
)

func main(){
  ctx := context.Background()
  options, err := redis.ParseURL("redis://localhost:6379")
  if err!=nil {
    log.Fatal(err)
  }
  redis_client := redis.NewClient(options)
  app := fiber.New()

  app.Get("/info", commands.Info(redis_client, ctx))
  app.Get("/get/:key", commands.Get(redis_client, ctx))

  app.Listen(":8080")
}
