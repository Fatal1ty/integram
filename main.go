// This entrypoint will run the single-process all-in-one instance of Integram
package main

import (
    "os",
	"github.com/requilence/integram"
	"github.com/integram-org/trello"
	"github.com/integram-org/gitlab"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	// You can fetch secrets from the ENV vars
	var trelloConfig trello.Config
	envconfig.MustProcess("TRELLO", &trelloConfig)

	integram.Register(
		trelloConfig,
		trelloConfig.BotConfig.Token,
	)

	// or just specify them directly
	integram.Register(
		gitlab.Config{
			OAuthProvider: integram.OAuthProvider{
				ID: os.GetEnv("GITLAB_APP_ID"),
				Secret: os.GetEnv("GITLAB_APP_SECRET"),
			},
		},
		os.GetEnv("TELEGRAM_BOT_TOKEN"),
	)
	integram.Run()

}