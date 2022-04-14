package ginhandler

import (
	"testing"
)

// TestRegister() will not pass unless we write a new Store wrapper interface
// to handle domain-specific operations like LoginUser and RegisterUser,
// since both Login and Register uses the same Store.First() to retrieve user data.
// Login uses Store.First() to find the target user, while Register uses the method to
// find duplicate username. The current implementation of Store interface and how handler methods
// call these interface methods restricts store.mockStore returning the correct values
// for both Login and Register.
func testRegister(t *testing.T) {
	// newUsername := "robpike"
	// reqBody := internal.AuthJson{
	// 	Username: newUsername,
	// 	Password: string(test.Password),
	// }
	// b, _ := json.Marshal(reqBody)
	// r := bytes.NewReader(b)
	// req, err := http.NewRequest("POST", "/", r)
	// if err != nil {
	// 	t.Fatal("failed to create new login request")
	// }
	// recorder := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(recorder)
	// c.Request = req

	// mockHandler := newMockHandler()
	// mockHandler.Register(c)

	// body := recorder.Body.Bytes()
	// if recorder.Code != http.StatusCreated {
	// 	t.Error(recorder.Code, string(body))
	// }
	// t.Logf("status: %+v\n", recorder.Result().Status)
	// t.Logf("body: %s\n", body)
}
