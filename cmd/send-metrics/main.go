package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
	"github.com/tyrannosaurus-becks/team-dashboard/internal"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	var config models.Config
	if err := envconfig.Process("dashboard", &config); err != nil {
		log.Fatal(err)
	}
	if err := internal.Run(&config); err != nil {
		log.Fatal(err)
	}
}
