package main

import (
	"log"
	"server/apis"
	"server/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	err := env.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(pprof.New())
	apis.Configure(app)
	app.Listen(env.Port)
}
