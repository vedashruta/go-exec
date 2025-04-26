package archive

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	obj "server/apis/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func create(c *fiber.Ctx) (err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(err)
	}
	files := form.File["file"]
	if err != nil {
		data := obj.Map{
			"ok":    false,
			"error": "no files uploaded",
		}
		return c.JSON(data)
	}
	if len(files) == 0 {
		data := obj.Map{
			"ok":    false,
			"error": "no files uploaded",
		}
		return c.JSON(data)
	}
	zipType := c.FormValue("type")
	var buf bytes.Buffer
	switch zipType {
	case "zip":
		err := createZip(files, &buf)
		if err != nil {
			data := obj.Map{
				"ok":    false,
				"error": err,
			}
			return c.JSON(data)
		}
		c.Set("Content-Type", "application/zip")
		c.Set("Content-Disposition", `attachment; filename="files.zip"`)
	case "tar.gz":
		err := createTarGz(files, &buf)
		if err != nil {
			data := obj.Map{
				"ok":    false,
				"error": err,
			}
			return c.JSON(data)
		}
		c.Set("Content-Type", "application/gzip")
		c.Set("Content-Disposition", `attachment; filename="files.tar.gz"`)
	default:
		data := obj.Map{
			"ok":    false,
			"error": fmt.Errorf("unsupported format: %s", zipType),
		}
		return c.JSON(data)
	}
	return c.SendStream(bytes.NewReader(buf.Bytes()))
}

func extract(c *fiber.Ctx) (err error) {
	file, err := c.FormFile("file")
	if err != nil {
		data := obj.Map{
			"ok":    false,
			"error": err,
		}
		return c.JSON(data)
	}
	buffer := bytes.Buffer{}
	multipartFile, err := file.Open()
	if err != nil {
		data := obj.Map{
			"ok":    false,
			"error": err,
		}
		return c.JSON(data)
	}
	_, err = buffer.ReadFrom(multipartFile)
	if err != nil {
		data := obj.Map{
			"ok":    false,
			"error": err,
		}
		return c.JSON(data)
	}
	base64 := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := fiber.Map{
		"ok":                       true,
		"base64 encoded multipart": base64,
	}
	return c.JSON(data)
}

func compress(c *fiber.Ctx) (err error) {
	type JSONRequest struct {
		Base64   string `json:"base64"`
		Filename string `json:"filename"`
		Alg      string `json:"alg"`
	}
	var (
		alg      string = "gzip"
		filename string = "file"
		data     io.Reader
	)
	cType := c.Get("Content-Type")
	switch cType {
	case fiber.MIMEApplicationJSON:
		var req JSONRequest
		err = json.Unmarshal(c.Body(), &req)
		if err != nil {
			data := fiber.Map{
				"ok":    false,
				"error": err,
			}
			return c.JSON(data)
		}
		if req.Base64 == "" {
			data := fiber.Map{
				"ok":    false,
				"error": "missing base64 data",
			}
			return c.JSON(data)
		}
		data = base64.NewDecoder(base64.StdEncoding, strings.NewReader(req.Base64))
		if req.Alg != "" {
			data := fiber.Map{
				"ok":    false,
				"error": "algorithm is required",
			}
			return c.JSON(data)
		}
		if req.Filename != "" {
			filename = req.Filename
		}
	default:
		fileHeader, err := c.FormFile("file")
		if err != nil {
			data := obj.Map{
				"ok":    false,
				"error": err,
			}
			return c.JSON(data)
		}
		f, err := fileHeader.Open()
		if err != nil {
			data := fiber.Map{
				"ok":    false,
				"error": err,
			}
			return c.JSON(data)
		}
		defer f.Close()
		data = f
		filename = fileHeader.Filename
	}
	c.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.%s"`, filename, alg))
	c.Set("Content-Type", "application/octet-stream")

	var writer io.WriteCloser
	switch alg {
	case "gzip":
		writer = gzip.NewWriter(c.Response().BodyWriter())
	case "zlib":
		writer = zlib.NewWriter(c.Response().BodyWriter())
	case "flate":
		var err error
		writer, err = flate.NewWriter(c.Response().BodyWriter(), flate.DefaultCompression)
		data := obj.Map{
			"ok":    false,
			"error": err,
		}
		return c.JSON(data)
	default:
		data := obj.Map{
			"ok":    false,
			"error": fmt.Errorf("unsupported compression algorithm"),
		}
		return c.JSON(data)
	}
	defer writer.Close()
	_, err = io.Copy(writer, data)
	if err != nil {
		data := obj.Map{
			"ok":    false,
			"error": err,
		}
		return c.JSON(data)
	}
	return
}
