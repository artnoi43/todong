package enums

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// ErrFile happens when values from *multipart.Form.File["foo"] fields an error.
	// When this is returned, a partial *multipartTodoData with nil FileData is returned
	ErrFile error = errors.New("error working with multipart file")
	// ErrPwTooShort happens when plaintext passwords are shorter than 6 characters
	ErrPwTooShort error = errors.New("password too short")
)

var (
	MapErrHandler = struct {
		JwtError       gin.H
		MultipartError gin.H
		Unmarshal      gin.H
	}{
		JwtError:       mapErrJwt,
		MultipartError: mapErrMultipart,
		Unmarshal:      mapErrUnmarshal,
	}

	mapErrJwt gin.H = gin.H{
		"status": "failed to extract jwt",
	}
	mapErrMultipart gin.H = gin.H{
		"status": "bad multipart request",
	}
	mapErrUnmarshal gin.H = gin.H{
		"status": "failed to unmarshal JSON",
	}
)
