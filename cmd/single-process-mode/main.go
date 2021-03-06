// This entrypoint will run the single-process all-in-one instance of Integram
package main

import (
    "os"
	"github.com/Fatal1ty/integram"
	// "github.com/integram-org/trello"
	"github.com/Fatal1ty/integram-gitlab"
	// "github.com/kelseyhightower/envconfig"
)

func main() {

	// You can fetch secrets from the ENV vars
	// var trelloConfig trello.Config
	// envconfig.MustProcess("TRELLO", &trelloConfig)

	// integram.Register(
	// 	trelloConfig,
	// 	trelloConfig.BotConfig.Token,
	// )

	// or just specify them directly
	integram.Register(
		gitlab.Config{
			OAuthProvider: integram.OAuthProvider{
				ID: os.Getenv("GITLAB_APP_ID"),
				Secret: os.Getenv("GITLAB_APP_SECRET"),
			},
		},
		os.Getenv("TELEGRAM_BOT_TOKEN"),
	)
	integram.Run()

}