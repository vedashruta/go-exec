package cpu

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/cpu")
	grp.Post("/memory", timeout.NewWithContext(memory, 5*time.Second))
	grp.Post("/system", timeout.NewWithContext(system, 5*time.Second))
}
