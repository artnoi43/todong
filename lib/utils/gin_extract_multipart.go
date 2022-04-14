package utils

import (
	"io/ioutil"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
)

// openAndReadMultipartFile opens and reads *multipart.FileHeader
func openAndReadMultipartFile(fh *multipart.FileHeader) ([]byte, error) {
	fp, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return ioutil.ReadAll(fp)
}

// ExtractTodoMultipartFileAndData extracts image file and JSON text data from multipart/form-data requests.
func ExtractTodoMultipartFileAndData(c *gin.Context) (*internal.MultipartTodoData, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get multipart form")
	}

	jsonData := form.Value["data"]
	// Only 1 multipart field with key "data" is supported
	if len(jsonData) != 1 {
		return nil, errors.New("only 1 json data field allowed")
	}
	jsonDataBytes := []byte(jsonData[0])

	// softFileError returns partial *internal.MultipartTodoData{}
	// The FileData field is set to nil since we likely encountered
	// some file error during read or encode (Base64) operations
	softFileError := func() (*internal.MultipartTodoData, error) {
		return &internal.MultipartTodoData{
			FileData: nil,
			JSONData: jsonDataBytes,
		}, enums.ErrFile
	}
	files := form.File["file"]
	// Only 1 multipart field with key "file" is supported
	if len(files) != 1 {
		return softFileError()
	}
	file := files[0]
	fileData, err := openAndReadMultipartFile(file)
	if err != nil {
		return softFileError()
	}
	return &internal.MultipartTodoData{
		FileData: fileData,
		JSONData: jsonDataBytes,
	}, nil
}
