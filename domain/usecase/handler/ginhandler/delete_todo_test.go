package ginhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/test"
)

func testDeleteTodo(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Params = []gin.Param{{Key: "uuid", Value: test.TodoUuid}}
	c.Request = postRequest
	testHandler := newTestHandler()
	testHandler.DeleteTodo(c)

	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusOK {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body: %s\n", body)
	var respBody = make(map[string]interface{})
	if err := json.Unmarshal(body, &respBody); err != nil {
		t.Error(fmt.Sprintf("failed to unmarshal json: %s", err.Error()))
	}
	if respBody["uuid"] != test.TodoUuid {
		t.Fatal("invalid uuid in response")
	}
}
