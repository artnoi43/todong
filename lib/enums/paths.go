package enums

type Endpoint string

var (
	HandlerRegister    Endpoint = "register"
	HandlerLogin       Endpoint = "login"
	HandlerCreateTodo  Endpoint = "create_todo"
	HandlerGetTodo     Endpoint = "get_todo"
	HandlerUpdateTodo  Endpoint = "update_todo"
	HandlerDeleteTodo  Endpoint = "delete_todo"
	HandlerNewPassword Endpoint = "new_password"
	HandlerDeleteUser  Endpoint = "delete_user"
	HandlerTestAuth    Endpoint = "test_auth"
)
