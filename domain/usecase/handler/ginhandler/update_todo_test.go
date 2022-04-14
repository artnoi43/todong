package ginhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/artnoi43/todong/test"
	"github.com/gin-gonic/gin"
)

func testUpdateTodo(t *testing.T) {
	req, _, jsonData, err := test.PrepareMultipartRequest()
	if err != nil {
		t.Fatalf("failed to create new multipart request: %s", err.Error())
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Request = req

	testHandler := newTestHandler()
	testHandler.UpdateTodo(c)

	if !bytes.Equal(test.TodoJson, jsonData) {
		t.Logf("expected: %s\n", test.TodoJson)
		t.Logf("actual: %s\n", jsonData)
		t.Fatal("multipart/form-data \"data\" not matched")
	}
	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusCreated {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body: %s\n", body)
	var respBody = make(map[string]interface{})
	if err := json.Unmarshal(body, &respBody); err != nil {
		t.Fatalf(fmt.Sprintf("failed to unmarshal json: %s", err.Error()))
	}
	if uuid := respBody["uuid"]; uuid != test.JwtIss {
		t.Logf("expected userUuid: %s\n", test.JwtIss)
		t.Logf("actual userUuid: %s\n", uuid)
		t.Fatal("invalid user uuid\n")
	}
}
