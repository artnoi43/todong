package gorillahandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/store"
	"github.com/artnoi43/todong/lib/utils"
)

func (h *GorillaHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]
	userUuid := r.Header.Get("iss")

	// To-do update request
	jsonData := strings.NewReader(r.FormValue("data"))
	respEncoder := json.NewEncoder(w)
	var updatesReq internal.TodoReqBody
	if err := json.NewDecoder(jsonData).Decode(&updatesReq); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(status)
		return
	}

	// To-do image
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to get file from request",
			"error":  err.Error(),
		})
		return
	}
	defer file.Close()
	imgBuf := new(bytes.Buffer)
	if _, err := io.Copy(imgBuf, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to read file",
			"error":  err.Error(),
		})
		return
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	if l := len(imgBase64Str); l > enums.POSTGRES_MAX_STRLEN {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"status":        "file too large",
			"file size":     l,
			"max file size": enums.POSTGRES_MAX_STRLEN,
		})
		return
	}

	// Get target to-do
	ctx := r.Context()
	var targetTodo datamodel.Todo
	if err := h.dataGateway.GetOneTodo(ctx, &datamodel.Todo{
		UUID:     uuid,
		UserUUID: userUuid,
	}, &targetTodo); err != nil {
		if errors.Is(err, store.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_ = respEncoder.Encode(map[string]interface{}{
				"status": "todo not found",
				"uuid":   uuid,
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to find todo",
			"uuid":   uuid,
			"error":  err.Error(),
		})
		return
	}

	// Generate new, updated todo
	var u datamodel.Todo
	utils.UpdatedTodo(uuid, imgBase64Str, &updatesReq, &targetTodo, &u)

	if err := h.dataGateway.UpdateTodo(ctx, &targetTodo, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed update todo",
			"error":  err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = respEncoder.Encode(map[string]interface{}{
		"status": "todo update successful",
		"uuid":   uuid,
	})
}
