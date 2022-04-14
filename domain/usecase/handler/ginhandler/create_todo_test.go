package ginhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/test"
)

func testCreateTodo(t *testing.T) {
	req, _, jsonData, err := test.PrepareMultipartRequest()
	if err != nil {
		t.Fatalf("failed to create new multipart request: %s", err.Error())
	}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Request = req

	testHandler := newTestHandler()
	testHandler.CreateTodo(c)

	if !bytes.Equal(test.TodoJson, jsonData) {
		t.Logf("expected: %s\n", test.TodoJson)
		t.Logf("actual: %s\n", jsonData)
		t.Fatal("multipart/form-data \"data\" not matched")
	}
	// if b := fileBuf.Bytes(); !bytes.Equal(b, formData.FileData) {
	// 	t.Logf("expected: %v\n", b)
	// 	t.Logf("actual: %v\n", formData.FileData)
	// 	t.Log("false alarm: multipart/form-data \"file\" not matched,\nalthough it works outside of this test")
	// }
	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusCreated {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body: %s\n", body)
	var newTodo model.Todo
	if err := json.Unmarshal(body, &newTodo); err != nil {
		t.Fatalf(fmt.Sprintf("failed to unmarshal json: %s", err.Error()))
	}
	if newTodo.UserUUID != test.JwtIss {
		t.Logf("expected userUuid: %s\n", test.JwtIss)
		t.Logf("actual userUuid: %s\n", newTodo.UserUUID)
		t.Fatal("invalid user uuid\n")
	}
	if newTodo.Status != enums.InProgress {
		t.Logf("expected userUuid: %s\n", enums.InProgress)
		t.Logf("actual userUuid: %s\n", newTodo.Status)
		t.Fatal("invalid todo status\n")
	}
	if newTodo.Description != test.TodoDescription {
		t.Logf("expected userUuid: %s\n", test.TodoDescription)
		t.Logf("actual userUuid: %s\n", newTodo.Description)
		t.Fatal("invalid todo description\n")
	}
	if newTodo.TodoDate != test.TodoDate {
		t.Logf("expected userUuid: %s\n", test.TodoDate)
		t.Logf("actual userUuid: %s\n", newTodo.TodoDate)
		t.Fatal("invalid todo date\n")
	}
}
