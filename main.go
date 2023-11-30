package main

import (
	"log"

	"github.com/iden3/go-service-template/config"
	"github.com/iden3/go-service-template/pkg/logger"
	httprouter "github.com/iden3/go-service-template/pkg/router/http"
	"github.com/iden3/go-service-template/pkg/router/http/handlers"
	"github.com/iden3/go-service-template/pkg/services/system"
	httptransport "github.com/iden3/go-service-template/pkg/transport/http"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	if err = logger.SetDefaultLogger(
		cfg.Log.Environment,
		cfg.Log.LogLevel(),
	); err != nil {
		log.Fatalf("failed to set default logger: %v", err)
	}

	err = newHTTPServic(cfg)
	if err != nil {
		logger.WithError(err).Error("failed to start http service")
	}
}

func newHTTPServic(cfg *config.Config) error {
	// init handlers
	systemHandlers := handlers.NewSystemHandler(
		system.NewReadinessService(),
		system.NewLivenessService(),
	)

	// init routers
	h := httprouter.NewHandlers(systemHandlers)
	routers := h.NewRouter(
		httprouter.WithOrigins(cfg.HTTPServer.Origins),
	)

	// run http server
	httpserver := httptransport.New(
		routers,
	)
	return httpserver.Start()
}
