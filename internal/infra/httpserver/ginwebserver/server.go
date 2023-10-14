package ginwebserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ismael3s/gowhatibought/internal/application/ports"
	"github.com/ismael3s/gowhatibought/internal/application/usecases"
)

type Server struct {
	engine            *gin.Engine
	findNfDataUseCase *usecases.FindNfDataUseCase
}

var _ ports.WebServer[gin.HandlerFunc] = &Server{}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) SetFindNfDataUseCase(useCase *usecases.FindNfDataUseCase) ports.WebServer[gin.HandlerFunc] {
	s.findNfDataUseCase = useCase
	return s
}

func (s *Server) RegisterHandler(method string, routePath string, handler gin.HandlerFunc) ports.WebServer[gin.HandlerFunc] {
	s.engine.Handle(method, routePath, handler)
	return s
}

func (s *Server) FindNfHandler(c *gin.Context) {
	result, err := s.findNfDataUseCase.Execute(usecases.FindNfDataUseCaseInput{
		Url: "http://nfe.sefaz.ba.gov.br/servicos/nfce/qrcode.aspx?p=29231004899316045480651010000305931390682954|2|1|1|c05f9c63156360dfa8dc4448f5f30ef244990eb8",
	})
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}

func FindNFHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong  -gin",
	})
}

func (s *Server) ListenAndServer(port string) error {
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
