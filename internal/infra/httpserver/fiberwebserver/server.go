package fiberwebserver

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ismael3s/gowhatibought/internal/application/ports"
	"github.com/ismael3s/gowhatibought/internal/application/usecases"
)

type Server struct {
	app               *fiber.App
	findNfDataUseCase *usecases.FindNfDataUseCase
}

var _ ports.WebServer[fiber.Handler] = &Server{}

func NewServer() *Server {
	return &Server{
		app: fiber.New(),
	}
}

func (s *Server) SetFindNfDataUseCase(useCase *usecases.FindNfDataUseCase) ports.WebServer[fiber.Handler] {
	s.findNfDataUseCase = useCase
	return s
}

func (s *Server) RegisterHandler(method string, routePath string, handler fiber.Handler) ports.WebServer[fiber.Handler] {
	s.app.Add(method, routePath, handler)
	return s
}

func FindNFHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong - foiber",
	})

}

func (s *Server) ListenAndServer(port string) error {
	return s.app.Listen(fmt.Sprintf(":%s", port))
}
