package fiberserver

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
}

func New() *fiberServer {
	return &fiberServer{
		app: fiber.New(),
	}
}

func (s *fiberServer) Shutdown(ctx context.Context) {
	if err := s.app.Shutdown(); err != nil {
		log.Printf("error shutting down fiber server: %s\n", err.Error())
	}
}
