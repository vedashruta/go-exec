package archive

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/archive")
	grp.Post("/create", timeout.NewWithContext(create, 5*time.Second))
	grp.Post("/extract", timeout.NewWithContext(extract, 5*time.Second))
	grp.Post("/compress", timeout.NewWithContext(compress, 5*time.Second))
	// grp.Post("/decompress", timeout.NewWithContext(decompress, 5*time.Second))
}
