package ginserver

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	engine *gin.Engine
}

func New() *ginServer {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return &ginServer{
		engine: r,
	}
}

// TODO: fix Shutdown panic for Gin
func (s *ginServer) Shutdown(ctx context.Context) {
	s.Shutdown(ctx)
}
