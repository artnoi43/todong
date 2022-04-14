package ginhandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/test"
)

// TODO: store.mockStore should returns valid bcrypt-hashed passwords
func testLogin(t *testing.T) {
	reqBody := internal.AuthJson{
		Username: test.Username,
		Password: string(test.Password),
	}
	b, _ := json.Marshal(reqBody)
	r := bytes.NewReader(b)
	req, err := http.NewRequest("POST", "/", r)
	if err != nil {
		t.Fatal("failed to create new login request")
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	testHandler := newTestHandler()
	testHandler.Login(c)

	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusAccepted {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body: %s\n", body)
}
