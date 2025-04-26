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
	app := fiber.New(
		fiber.Config{
			StrictRouting: true,
			CaseSensitive: true,
			BodyLimit:     env.BodyLimit,
			Concurrency:   env.Concurrency,
		},
	)
	app.Use(pprof.New())
	app.Get("/", func(c *fiber.Ctx) (err error) {
		c.JSON(fiber.Map{
			"ok":     false,
			"status": "denied",
		})
		return
	})
	apis.Configure(app)
	app.Listen(env.Port)
}
