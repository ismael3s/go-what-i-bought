package main

import (
	"os"

	"github.com/ismael3s/gowhatibought/internal/application/ports"
	"github.com/ismael3s/gowhatibought/internal/infra/di"
)

func main() {
	webServerPort := os.Getenv("PORT")
	if webServerPort == "" {
		webServerPort = "8080"
	}
	runServer(di.NewWebServerFiber(), webServerPort)
}

func runServer(server ports.WebServerRunner, port string) error {
	return server.ListenAndServer(port)
}
