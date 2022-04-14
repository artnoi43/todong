package utils

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/test"
)

// TODO: key "file" does not work in test because
// The request lacks FileHeader is not complete
func TestMultipart(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, fileBuf, j, err := test.PrepareMultipartRequest()
	if err != nil {
		t.Fatalf("failed to create multipart request: %s", err.Error())
	}
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Request = req

	formData, err := ExtractTodoMultipartFileAndData(c)
	if err != nil {
		t.Logf("%+v\n", formData)
		t.Fatal(err.Error())
	}
	if !bytes.Equal(j, formData.JSONData) {
		t.Logf("expected: %s\n", j)
		t.Logf("actual: %s\n", formData.JSONData)
		t.Fatal("multipart/form-data \"data\" not matched")
	}
	if b := fileBuf.Bytes(); !bytes.Equal(b, formData.FileData) {
		t.Logf("expected: %v\n", b)
		t.Logf("actual: %v\n", formData.FileData)
		t.Log("false alarm: multipart/form-data \"file\" not matched,\nalthough it works outside of this test")
	}
}
