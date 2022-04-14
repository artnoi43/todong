package ginhandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/test"
	"github.com/gin-gonic/gin"
)

func testNewPassword(t *testing.T) {
	reqBody := internal.NewPasswordJson{
		NewPassword: string(test.Password),
	}
	b, _ := json.Marshal(reqBody)
	r := bytes.NewReader(b)
	req, err := http.NewRequest("POST", "/", r)
	if err != nil {
		t.Fatal("failed to create new login request")
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Request = req

	testHandler := newTestHandler()
	testHandler.NewPassword(c)

	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusOK {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body: %s\n", body)
}
