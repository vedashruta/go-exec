package json

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Map map[string]any

func encode(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	resBytes, err := Encode(string(bytes))
	if err != nil {
		return c.JSON(err)
	}
	c.Set("Content-Type", "application/octet-stream")
	data := Map{
		"data in bytes": resBytes,
	}
	return c.JSON(data)
}

func Encode(input any) (result []byte, err error) {
	result, err = json.Marshal(input)
	if err != nil {
		return
	}
	return
}

func Decode(input []byte, toAddress any) (err error) {
	err = json.Unmarshal(input, &toAddress)
	if err != nil {
		return
	}
	return
}
