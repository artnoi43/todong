package utils

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/internal"
)

func ExtractTodoMultipartFileAndDataFiber(c *fiber.Ctx) (*internal.MultipartTodoData, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get multipart/form-data field \"file\"")
	}
	fp, err := file.Open() // No need to Close()
	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, fp); err != nil {
		return nil, errors.Wrap(err, "failed to read multipart/form-data file")
	}
	data := c.FormValue("data")
	if data == "" {
		return nil, errors.Wrap(err, "failed to get multipart/form-data field \"data\"")
	}
	return &internal.MultipartTodoData{
		FileData: fileBuf.Bytes(),
		JSONData: []byte(data),
	}, nil
}
