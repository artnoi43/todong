package gorillahandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/lib/store"
)

func (h *GorillaHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]
	userUuid := r.Header.Get("iss")

	respEncoder := json.NewEncoder(w)
	ctx := r.Context()
	if err := h.dataGateway.DeleteTodo(ctx, &datamodel.Todo{
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
