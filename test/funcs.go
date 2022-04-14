package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
)

var imagePath = "test.txt"

func PrepareMultipartRequest() (*http.Request, *bytes.Buffer, []byte, error) {
	fp, err := os.Open(imagePath)
	if err != nil {
		return nil, nil, nil, err
	}
	fi, _ := os.Stat(imagePath)
	// For comparison
	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, fp); err != nil {
		return nil, nil, nil, err
	}
	defer fp.Close()
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	// Create the form data field fp
	dataPart, err := mw.CreateFormField("data")
	if err != nil {
		return nil, nil, nil, err
	}
	reqBody := Todo
	j, _ := json.Marshal(reqBody)
	jr := bytes.NewReader(j)
	if _, err := io.Copy(dataPart, jr); err != nil {
		return nil, nil, nil, err
	}

	// Create the form data field 'file'
	filePart, err := mw.CreateFormFile("file", fi.Name())
	if err != nil {
		return nil, nil, nil, err
	}
	if _, err := io.Copy(filePart, fp); err != nil {
		return nil, nil, nil, err
	}

	mw.Close()
	// Our test HTTP request
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())
	return req, fileBuf, j, nil
}
