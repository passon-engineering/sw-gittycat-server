package processor

import (
	"fmt"
	"sw-gittycat-server/modules/actions"
	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/git"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func Init(app *application.Application) {
	go func() {
		for action := range app.Queue {
			processWebhookQueue(action, app)
		}
	}()
}

func processWebhookQueue(action actions.Action, app *application.Application) {
	startTime := time.Now()

	git.PullRepo(action.Webhook.RepoName, app)
	commandHandler := func(command string, output string, err error) {
		if err != nil {
			output := fmt.Sprintf("'%s': Command '%s' failed: %v", action.Webhook.RepoName, command, err)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Error:          output,
				ProcessingTime: time.Since(startTime),
			})
		} else {
			output := fmt.Sprintf("'%s': Command '%s' executed successfully, output: %s", action.Webhook.RepoName, command, output)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           output,
				ProcessingTime: time.Since(startTime),
			})
		}
	}

	action.Webhook.RunCommands(commandHandler)
}
