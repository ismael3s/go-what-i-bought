package di

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/ismael3s/gowhatibought/internal/application/ports"
	"github.com/ismael3s/gowhatibought/internal/application/usecases"
	"github.com/ismael3s/gowhatibought/internal/infra/adapters"
	"github.com/ismael3s/gowhatibought/internal/infra/httpserver/fiberwebserver"
	"github.com/ismael3s/gowhatibought/internal/infra/httpserver/ginwebserver"
)

func NewFindNfDataUsecase() *usecases.FindNfDataUseCase {
	fazendaGateway := adapters.NewFazendaGateway()
	brasilApiGateway := adapters.NewBrasilApiGateway()
	return usecases.NewFindNfDataUsecase(fazendaGateway, brasilApiGateway)
}

func NewWebServer() ports.WebServerRunner {
	server := ginwebserver.NewServer()
	server.SetFindNfDataUseCase(NewFindNfDataUsecase()).
		RegisterHandler("GET", "/api/v1/find-nf", server.FindNfHandler).
		RegisterHandler("GET", "/api/v1/health-check", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "health",
			})
		})
	return server
}

func NewWebServerFiber() ports.WebServerRunner {
	server := fiberwebserver.NewServer()
	server.
		SetFindNfDataUseCase(NewFindNfDataUsecase()).
		RegisterHandler("GET", "/api/v1/find-nf", server.FindNfHandler).
		RegisterHandler("GET", "/api/v1/health-check", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "health",
			})
		})
	return server
}
