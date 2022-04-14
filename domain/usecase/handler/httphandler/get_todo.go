package httphandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/domain/usecase"
)

func (h *HttpHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	var uuid string
	if r.Context().Value(usecase.ContextKeyNetHttp) != nil {
		p := r.URL.Path
		uuid = strings.TrimPrefix(p, "/")
	} else {
		params := mux.Vars(r)
		uuid = params["uuid"]
	}
	userUuid := r.Header.Get("iss")

	var getAll bool
	if len(uuid) == 0 {
		getAll = true
	}
	respEncoder := json.NewEncoder(w)
	ctx := r.Context()
	var todos []model.Todo
	if getAll {
		if err := h.dataGateway.GetUserTodos(ctx, &model.Todo{
			UserUUID: userUuid,
		}, &todos); err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
				_ = respEncoder.Encode(map[string]interface{}{
					"status":   "todos not found",
					"userUuid": userUuid,
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			_ = respEncoder.Encode(map[string]interface{}{
				"status": "failed to find todo",
				"uuid":   uuid,
				"error":  err.Error(),
			})
			return
		}
	} else {
		var todo model.Todo
		if err := h.dataGateway.GetOneTodo(ctx, &model.Todo{
			UserUUID: userUuid,
			UUID:     uuid,
		}, &todo); err != nil {
			if errors.Is(err, store.ErrRecordNotFound) {
				w.WriteHeader(http.StatusNotFound)
				_ = respEncoder.Encode(map[string]interface{}{
					"status": "todos not found",
					"uuid":   uuid,
					"error":  err.Error(),
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			_ = respEncoder.Encode(map[string]interface{}{
				"status": "failed to find todo",
				"uuid":   uuid,
				"error":  err.Error(),
			})
			return
		}
		todos = append(todos, todo)
	}
	if len(todos) == 0 {
		w.WriteHeader(http.StatusNotFound)
		respEncoder.Encode(map[string]interface{}{
			"status": "todos not found",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = respEncoder.Encode(todos)
}
