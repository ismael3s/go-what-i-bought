package fiberwebserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ismael3s/gowhatibought/internal/application/usecases"
)

func (s *Server) FindNfHandler(c *fiber.Ctx) error {
	url := c.Query("url")
	result, err := s.findNfDataUseCase.Execute(usecases.FindNfDataUseCaseInput{
		Url: url,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(result)
}
