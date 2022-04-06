package gorillahandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/lib/store"
)

func (h *GorillaHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userUuid := r.Header.Get("iss")
	respEncoder := json.NewEncoder(w)

	ctx := r.Context()
	var targetUser datamodel.User
	if err := h.dataGateway.GetUserByUuid(ctx, userUuid, &targetUser); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_ = respEncoder.Encode(map[string]interface{}{
				"status": "user not found",
				"uuid":   userUuid,
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to get target user",
			"uuid":   userUuid,
			"error":  err.Error(),
		})
		return
	}
	if err := h.dataGateway.DeleteUser(ctx, &targetUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to delete target user",
			"uuid":   userUuid,
			"error":  err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	_ = respEncoder.Encode(map[string]interface{}{
		"status": "user deletion successful",
		"uuid":   userUuid,
	})
}
