package gorillahandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
)

func (h *GorillaHandler) NewPassword(w http.ResponseWriter, r *http.Request) {
	userUuid := r.Header.Get("iss")
	respEncoder := json.NewEncoder(w)

	var req internal.NewPasswordJson
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(status)
		return
	}
	if len(req.NewPassword) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"error": "blank password",
			"uuid":  userUuid,
		})
		return
	}
	pw, err := utils.EncodeBcrypt([]byte(req.NewPassword))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to get target user",
			"uuid":   userUuid,
		})
		return
	}

	ctx := r.Context()
	var targetUser datamodel.User
	if err := h.dataGateway.GetUserByUuid(ctx, userUuid, &targetUser); err != nil {
		// Should not happen, but possible
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
		})
		return
	}

	if err := h.dataGateway.ChangePassword(ctx, &targetUser, pw); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to update password",
			"uuid":   userUuid,
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = respEncoder.Encode(map[string]interface{}{
		"status": "password update successful",
		"uuid":   userUuid,
	})
}
