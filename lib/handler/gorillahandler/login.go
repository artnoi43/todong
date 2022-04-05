package gorillahandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
)

func (h *GorillaHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req internal.AuthJson
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respEncoder := json.NewEncoder(w)
	var user datamodel.User
	ctx := r.Context()
	if err := h.dataGateway.GetUserByUsername(ctx, req.Username, &user); err != nil {
		if err != store.ErrRecordNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			respEncoder.Encode(map[string]interface{}{
				"error": "login failed",
			})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"error": "invalid username or password",
		})
		return
	}
	if user.UUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"error": "invalid username or password",
		})
		return
	}
	if err := utils.DecodeBcrypt(user.Password, []byte(req.Password)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respEncoder.Encode(map[string]interface{}{
			"error": "invalid username or password",
		})
		return
	}
	token, exp, err := utils.NewJwtToken(user.UUID, []byte(h.config.SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respEncoder.Encode(map[string]interface{}{
			"error": "login failed",
		})
		return
	}
	resp := internal.LoginResponse(struct {
		Status   string
		Username string
		UserUuid string
		Exp      time.Time
		Token    string
	}{
		Status:   "login successful",
		Username: user.Username,
		UserUuid: user.UUID,
		Exp:      exp,
		Token:    token,
	})
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Marshal())
}
