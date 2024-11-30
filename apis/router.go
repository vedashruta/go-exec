package apis

import (
	"server/apis/base64"
	"server/apis/callback"
	"server/apis/cpu"
	"server/apis/cryptography"
	"server/apis/json"
	"server/apis/jwt"

	"github.com/gofiber/fiber/v2"
)

func Configure(app *fiber.App) {
	grp := app.Group("/api")
	route(grp)
}
func route(r fiber.Router) {
	callback.Route(r)
	json.Route(r)
	base64.Route(r)
	cpu.Route(r)
	jwt.Route(r)
	cryptography.Route(r)
}
