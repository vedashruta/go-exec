package base64

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"server/apis/json"
	"server/env"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	Text     string `json:"text,omitempty"`
	Base64   string `json:"base64,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func encodeMultipart(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	base64, err := base64.StdEncoding.DecodeString(req.Base64)
	if err != nil {
		return c.JSON(err)
	}
	path := fmt.Sprintf("%[1]s/%[2]s", env.Downloads, req.FileName)
	f, err := os.Create(path)
	if err != nil {
		return c.JSON(err)
	}
	defer os.Remove(path)
	if _, err := f.Write(base64); err != nil {
		return c.JSON(err)
	}
	return c.SendFile(path)
}

func decodeMultipart(c *fiber.Ctx) (err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(err)
	}
	buffer := bytes.Buffer{}
	multipartFile, err := file.Open()
	if err != nil {
		return c.JSON(err)
	}
	_, err = buffer.ReadFrom(multipartFile)
	if err != nil {
		return c.JSON(err)
	}
	base64 := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := fiber.Map{
		"base64 encoded multipart": base64,
	}
	return c.JSON(data)

}
func encode(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	base64 := base64.StdEncoding.EncodeToString([]byte(req.Text))
	data := json.Map{
		"base64 encoded string": base64,
	}
	return c.JSON(data)
}

func decode(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	text, err := base64.StdEncoding.DecodeString(req.Base64)
	if err != nil {
		return c.JSON(err)
	}
	data := json.Map{
		"base64 decoded string": string(text),
	}
	return c.JSON(data)
}

func urlEncode(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	text := base64.URLEncoding.EncodeToString([]byte(req.Text))
	data := json.Map{
		"base64 url encoded string": string(text),
	}
	return c.JSON(data)
}

func urlDecode(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	type request struct {
		Base64 string `json:"base64"`
	}
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	stringBytes, err := base64.URLEncoding.DecodeString(req.Base64)
	if err != nil {
		return c.JSON(err)
	}
	data := json.Map{
		"base64 url decoded string": string(stringBytes),
	}
	return c.JSON(data)
}
