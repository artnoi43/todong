package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/utils"
)

func (h *HttpHandler) Register(w http.ResponseWriter, r *http.Request) {
	respEncoder := json.NewEncoder(w)
	var req internal.AuthJson
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(status)
		return
	}

	// Also, return after calling this func
	registerFailed := func() {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"error": "register failed",
		})
	}

	var user model.User
	ctx := r.Context()
	_ = h.dataGateway.GetUserByUsername(ctx, req.Username, &user)

	if user.Username != "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"duplicate username": req.Username,
		})
		return
	}
	pw, err := utils.EncodeBcrypt([]byte(req.Password))
	if err != nil {
		registerFailed()
		return
	}
	user = model.User{
		UUID:     uuid.NewString(),
		Username: req.Username,
		Password: pw,
	}
	if err := h.dataGateway.CreateUser(ctx, &user); err != nil {
		registerFailed()
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = respEncoder.Encode(map[string]interface{}{
		"username": user.Username,
		"userUuid": user.UUID,
	})
}
