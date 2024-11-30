package base64

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(r fiber.Router) {
	grp := r.Group("/base64")
	grp.Post("/encodemultipart", timeout.NewWithContext(encodeMultipart, 5*time.Second))
	grp.Post("/decodemultipart", timeout.NewWithContext(decodeMultipart, 5*time.Second))
	grp.Post("/encode", timeout.NewWithContext(encode, 5*time.Second))
	grp.Post("/decode", timeout.NewWithContext(decode, 5*time.Second))
	grp.Post("/urlencode", timeout.NewWithContext(urlEncode, 5*time.Second))
	grp.Post("/urldecode", timeout.NewWithContext(urlDecode, 5*time.Second))
}
