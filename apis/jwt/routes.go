package jwt

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/jwt")
	grp.Post("/generate", timeout.NewWithContext(generate, 5*time.Second))
	grp.Post("/validate", timeout.NewWithContext(validate, 5*time.Second))
}
