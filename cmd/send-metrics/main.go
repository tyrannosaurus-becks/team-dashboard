package main

import (
	"log"
	"os"

	"github.com/tyrannosaurus-becks/team-dashboard/internal"
	"github.com/tyrannosaurus-becks/team-dashboard/internal/models"
)

func main() {
	if err := internal.Run(&models.Config{
		DatadogClientAPIKey: os.Getenv("DD_CLIENT_API_KEY"),
		DatadogClientAppKey: os.Getenv("DD_CLIENT_APP_KEY"),
		GithubAccessToken:   os.Getenv("GITHUB_ACCESS_TOKEN"),
	}); err != nil {
		log.Fatal(err)
	}
}
