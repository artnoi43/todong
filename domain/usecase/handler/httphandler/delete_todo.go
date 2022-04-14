package httphandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/data/store"
	"github.com/artnoi43/todong/domain/usecase"
)

func (h *HttpHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var uuid string
	if r.Context().Value(usecase.ContextKeyNetHttp) != nil {
		p := r.URL.Path
		uuid = strings.TrimPrefix(p, "/")
	} else {
		params := mux.Vars(r)
		uuid = params["uuid"]
	}
	userUuid := r.Header.Get("iss")
	fmt.Println(userUuid, uuid)

	respEncoder := json.NewEncoder(w)
	ctx := r.Context()
	if err := h.dataGateway.DeleteTodo(ctx, &model.Todo{
		UUID:     uuid,
		UserUUID: userUuid,
	}); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_ = respEncoder.Encode(map[string]interface{}{
				"status": "todo not found",
				"uuid":   uuid,
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to delete todo",
			"uuid":   uuid,
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = respEncoder.Encode(map[string]interface{}{
		"status": "todo delete successful",
		"uuid":   uuid,
	})
}
