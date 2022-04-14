package ginhandler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/test"
)

var (
	postRequest = httptest.NewRequest("POST", "/cdb1eb67-fe8c-44ba-a32a-bda8ca95bd3c", nil)
)

func testGetTodo(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Params = []gin.Param{{Key: "uuid", Value: "cdb1eb67-fe8c-44ba-a32a-bda8ca95bd3c"}}
	c.Request = postRequest
	testHandler := newTestHandler()
	testHandler.GetTodo(c)

	if recorder.Code != http.StatusOK {
		b, _ := ioutil.ReadAll(recorder.Body)
		t.Error(recorder.Code, string(b))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body (expecting null, because using mocked data store): %s\n", recorder.Body.Bytes())
}

func testGetTodos(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Keys = test.JwtKeys
	c.Request = postRequest
	testHandler := newTestHandler()
	testHandler.GetTodo(c)

	body := recorder.Body.Bytes()
	if recorder.Code != http.StatusOK {
		t.Error(recorder.Code, string(body))
	}
	t.Logf("status: %+v\n", recorder.Result().Status)
	t.Logf("body (expecting null, because using mocked data store): %s\n", recorder.Body.Bytes())
}

// func TestGetTodoNew(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockDataGateway := mockstore.NewMockDataGateway(ctrl)
// 	_ = mockDataGateway
// }
