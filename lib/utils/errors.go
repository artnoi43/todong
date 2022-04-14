package utils

import "github.com/gin-gonic/gin"

func ErrStatus(m gin.H, err error) map[string]interface{} {
	m["error"] = err.Error()
	return m
}
