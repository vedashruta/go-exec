package cryptography

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"server/apis/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	Algorithm string `json:"algorithm"`
	PlainText string `json:"plain_text"`
	Key       string `json:"key,omitempty"`
}

func GenerateRandomBytes(length int) (bytes []byte, err error) {
	bytes = make([]byte, length)
	_, err = io.ReadFull(rand.Reader, bytes)
	if err != nil {
		return
	}
	return
}

func generateKeys(c *fiber.Ctx) (err error) {
	reqBytes := c.Request().Body()
	req := request{}
	err = json.Decode(reqBytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	var cryptography Cryptography
	switch strings.ToLower(req.Algorithm) {
	case "aes":
		cryptography = &AES{}
	case "des":
		cryptography = &DES{}
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported algorithm"})
	}
	key, err := cryptography.GenerateKey()
	if err != nil {
		return c.JSON(err)
	}
	encodedKey := base64.StdEncoding.EncodeToString([]byte(key))
	data := json.Map{
		"encoded key": encodedKey,
	}
	return c.JSON(data)
}

func encrypt(c *fiber.Ctx) (err error) {
	reqBytes := c.Request().Body()
	req := request{}
	err = json.Decode(reqBytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	var cryptography Cryptography
	switch strings.ToLower(req.Algorithm) {
	case "aes":
		cryptography = &AES{
			Key:       req.Key,
			PlainText: req.PlainText,
		}
	case "des":
		cryptography = &DES{
			Key:       req.Key,
			PlainText: req.PlainText,
		}
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported algorithm"})
	}
	cipherText, err := cryptography.Encrypt()
	if err != nil {
		return c.JSON(err)
	}
	data := json.Map{
		"cipher": cipherText,
	}
	return c.JSON(data)
}

func decrypt(c *fiber.Ctx) (err error) {
	return
}

func hash(c *fiber.Ctx) (err error) {
	return
}
