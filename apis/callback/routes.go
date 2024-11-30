package callback

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/callback")
	grp.Post("/callback1", timeout.NewWithContext(callback1, 5*time.Second))
	grp.Post("/callback2", timeout.NewWithContext(callback2, 5*time.Second))
	grp.Post("/multipart", timeout.NewWithContext(multipart, 5*time.Second))
}
