// cmd/migrations/main.go
package main

import (
	"os"

	"github.com/lob/logger-go"
	"github.com/robinjoseph08/go-pg-migrations"
	"github.com/sachinmurali/goyagi/pkg/config"
	"github.com/sachinmurali/goyagi/pkg/database"
)

const directory = "cmd/migrations"

func main() {
	log := logger.New()

	cfg := config.New()

	db, err := database.New(cfg)
	if err != nil {
		log.Err(err).Fatal("failed to connect to database")
	}

	err = migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Err(err).Fatal("failed migration")
	}
}
