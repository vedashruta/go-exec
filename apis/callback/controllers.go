package callback

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func callback1(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	fmt.Println("Callback 1 is called")
	fmt.Println("payload: ", string(bytes))
	return
}

func callback2(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	fmt.Println("Callback 2 is called")
	fmt.Println("payload: ", string(bytes))
	return
}

func multipart(c *fiber.Ctx) (err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return c.JSON(fiber.Map{"err": "failed to determine user's home directory"})
	}
	path := filepath.Join(homeDir, "Downloads", file.Filename)
	err = c.SaveFile(file, path)
	if err != nil {
		return c.JSON(err)
	}
	return
}
