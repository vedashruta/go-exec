package json

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/json")
	grp.Post("/encode", timeout.NewWithContext(encode, 5*time.Second))
}
