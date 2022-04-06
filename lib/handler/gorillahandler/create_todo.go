package gorillahandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/artnoi43/todong/datamodel"
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/utils"
)

func (h *GorillaHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	respEncoder := json.NewEncoder(w)
	file, _, err := r.FormFile("file")
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.MultipartError, err)
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(status)
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
	jsonBody := r.FormValue("data")
	if len(jsonBody) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "empty multipart/form-data key \"data\"",
		})
		return
	}

	var req internal.TodoReqBody
	if err := json.Unmarshal([]byte(jsonBody), &req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		w.WriteHeader(http.StatusBadRequest)
		_ = respEncoder.Encode(status)
		return
	}
	userUuid := r.Header.Get("iss")
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	order, err := datamodel.NewTodo(userUuid, req.Title, req.Description, req.TodoDate, enums.Status(req.Status), imgBase64Str)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to create new todo (1)",
			"error":  err.Error(),
		})
		return
	}
	if err := h.dataGateway.CreateTodo(r.Context(), order); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = respEncoder.Encode(map[string]interface{}{
			"status": "failed to create new todo (2)",
			"error":  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = respEncoder.Encode(order)
}
