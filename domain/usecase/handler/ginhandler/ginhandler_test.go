package ginhandler

import (
	"testing"

	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/domain/usecase/middleware"
)

// Use our manually written store.testDataGateway
func newTestHandler() *GinHandler {
	return &GinHandler{
		dataGateway: store.NewTestDataGateway(),
		config: &middleware.Config{
			SecretKey: "qwerty",
		},
	}
}

func TestHandler(t *testing.T) {
	t.Run("Test Register", testRegister) // empty
	t.Run("Test Login", testLogin)
	t.Run("Test NewPassword", testNewPassword)
	t.Run("Test GetTodo", testGetTodo)
	t.Run("Test GetTodos", testGetTodos)
	t.Run("Test CreateTodo", testCreateTodo)
	t.Run("Test DeleteTodo", testDeleteTodo)
	t.Run("Test DeleteUser", testDeleteUser)
	t.Run("Test UpdateTodo", testUpdateTodo)
}
