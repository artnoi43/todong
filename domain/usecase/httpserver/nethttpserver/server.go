package nethttpserver

import (
	"context"
	"net/http"

	"github.com/artnoi43/todong/domain/usecase"
	"github.com/artnoi43/todong/domain/usecase/handler"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

type nethttpServer struct {
	adapter handler.Adapter
	router  *http.ServeMux
	auth    *utils.Authenticator
}

func New() *nethttpServer {
	// adapter will be initialized by SetUpRoutes()
	return &nethttpServer{
		router: http.NewServeMux(),
	}
}

func (h *nethttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	// /todos/kuy => todo /kuy
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	switch head {
	case "users":
		h.handleUsers(w, r)
		return
	case "todos":
		h.handleTodos(w, r)
		return
	}
	http.NotFound(w, r)
}

func (h *nethttpServer) handleUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	var head string
	head, _ = utils.ShiftPath(r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		switch head {
		case "register":
			h.adapter.NetHttp(enums.HandlerRegister)(w, r)
		case "login":
			h.adapter.NetHttp(enums.HandlerLogin)(w, r)
		}
	case http.MethodDelete:
		h.authRequired(enums.HandlerDeleteUser, w, r)
	}
}

func (s *nethttpServer) handleTodos(
	w http.ResponseWriter,
	r *http.Request,
) {
	var head string
	head, _ = utils.ShiftPath(r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "":
			fallthrough
		default:
			s.authRequired(enums.HandlerGetTodo, w, r)
		}
	case http.MethodPost:
		switch head {
		case "create":
			s.authRequired(enums.HandlerCreateTodo, w, r)
		case "update":
			s.authRequired(enums.HandlerUpdateTodo, w, r)
		}
	case http.MethodDelete:
		s.authRequired(enums.HandlerDeleteTodo, w, r)
	}
}

func (s *nethttpServer) wrapMw(next http.Handler) http.Handler {
	return s.auth.AuthMiddleware(next)
}

func (s *nethttpServer) authRequired(t enums.Endpoint, w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), usecase.ContextKeyNetHttp, true)
	var _handler http.Handler
	switch t {
	case enums.HandlerLogin, enums.HandlerRegister:
		_handler = s.adapter.NetHttp(t)
	default:
		_handler = s.wrapMw(s.adapter.NetHttp(t))
	}
	_handler.ServeHTTP(w, r.WithContext(ctx))
}
