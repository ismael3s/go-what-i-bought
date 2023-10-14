package ports

type WebServer[handlerType any] interface {
	RegisterHandler(method string, routePath string, handler handlerType) WebServer[handlerType]
}

type WebServerRunner interface {
	ListenAndServer(port string) error
}
