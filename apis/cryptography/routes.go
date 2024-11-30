package cryptography

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/cryptography")
	grp.Post("/encrypt", timeout.NewWithContext(encrypt, 5*time.Second))
	// grp.Post("/decrypt", timeout.NewWithContext(decrypt, 5*time.Second))
	grp.Post("/generatekeys", timeout.NewWithContext(generateKeys, 5*time.Second))
}
