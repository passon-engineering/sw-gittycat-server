package processor

import (
	"fmt"
	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/git"
	"sw-gittycat-server/modules/webhooks"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func Init(app *application.Application) {
	go func() {
		fmt.Println("TEST")
		for action := range app.Queue {
			processWebhookQueue(action, app)
		}
	}()
}

func processWebhookQueue(action webhooks.WebhookAction, app *application.Application) {
	startTime := time.Now()

	switch action.Action {
	case "delete":
		// handle add
		// Possible future implementation placeholder
		return
	case "add":
		// handle add
		// Possible future implementation placeholder
		return
	case "run":
		// handle run
		git.PullRepo(action.Webhook.RepoName, app)
		commandHandler := func(command string, output string, err error) {
			if err != nil {
				output := fmt.Sprintf("'%s': Command '%s' failed: %v", action.Webhook.RepoName, command, err)
				app.Logger.Entry(logger.Container{
					Status: logger.STATUS_ERROR,
					Error:  output,
				})
			} else {
				output := fmt.Sprintf("'%s': Command '%s' executed successfully, output: %s", action.Webhook.RepoName, command, output)
				app.Logger.Entry(logger.Container{
					Status: logger.STATUS_INFO,
					Info:   output,
				})
			}
		}

		action.Webhook.RunCommands(commandHandler)
		app.WebhookHandler.UpdateLastCall(action.Webhook.RepoName, string(startTime.Format(time.RFC3339)))
	default:
		fmt.Println("Action not found " + action.Action)
	}

	fmt.Printf("Processing webhook action: repo = %s, action = %s\n", action.Webhook.RepoName, action.Action)
}
