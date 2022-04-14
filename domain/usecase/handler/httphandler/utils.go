package httphandler

import "net/http"

// type CtxKey struct{}

// Maybe useful for net/http server implementation
func GetField(r *http.Request) []string {
	fields := r.Context().Value("ctxKey").([]string)
	return fields
}
