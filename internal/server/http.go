package server

import (
	v1 "blog/api/helloworld/v1"
	"blog/configs"
	"blog/internal/service"
	nethttp "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *configs.ApplicationConfig, greeter *service.GreeterService, logger log.Logger) *http.Server {
	// func NewHTTPServer(c *configs.ApplicationConfig, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Server.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.Server.HTTP.Addr))
	}
	if c.Server.HTTP.Timeout != 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Server.HTTP.Timeout)))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	openAPIhandler := handleSwaggerUI(configs.OpenAPI)
	srv.HandlePrefix("/q/", openAPIhandler)
	return srv
}

func handleSwaggerUI(file []byte) nethttp.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/q/openapi.yaml", byteHandler(file))
	return router
}

func byteHandler(b []byte) nethttp.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(b)
	}
}
