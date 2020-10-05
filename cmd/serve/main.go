// cmd/serve/main.go
package main

import (
	"net/http"

	"github.com/lob/logger-go"
	"github.com/sachinmurali/goyagi/pkg/application"
	"github.com/sachinmurali/goyagi/pkg/server"
)

func main() {
	log := logger.New()

	app, err := application.New()
	if err != nil {
		log.Err(err).Fatal("failed to initialize application")
	}

	srv := server.New(app)

	log.Info("server started", logger.Data{"port": app.Config.Port})

	errors := srv.ListenAndServe()
	if errors != nil && errors != http.ErrServerClosed {
		log.Err(errors).Fatal("server stopped")
	}

	log.Info("server stopped")
}
