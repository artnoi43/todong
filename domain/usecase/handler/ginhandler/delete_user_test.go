package ginhandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/test"
)

func testDeleteUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Request, _ = http.NewRequest("POST", "/", nil)

	testHandler := newTestHandler()
	testHandler.DeleteUser(c)

	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusOK {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body: %s\n", body)
	var resp = make(map[string]interface{})
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Errorf("error unmarshaling json response: %s\n", err.Error())
	}
	if userUuid := resp["uuid"]; userUuid != test.JwtIss {
		t.Logf("expected userUuid: %s\n", test.JwtIss)
		t.Logf("actual userUuid: %s\n", userUuid)
		t.Fatal("invalid userUuid\n")
	}
}
