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
			processActionQueue(action, app)
		}
	}()
}

func processActionQueue(action actions.Action, app *application.Application) {
	startTime := time.Now()

	var success = true

	git.PullRepo(action.Webhook.RepoName, app)
	commandHandler := func(command string, output string, err error) {
		if err != nil {
			success = false
			output := fmt.Sprintf("'%s': Command '%s' failed: %v", action.Webhook.RepoName, command, err)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Error:          output,
				ProcessingTime: time.Since(startTime),
			})
		} else {
			after, ok := action.RequestBody["after"].(string)
			if !ok || len(after) == 0 {
				after = "no_hash_available"
			}

			elapsedTime := time.Since(startTime)
			action.ProcessingTime = elapsedTime.String()

			action.FileName = startTime.Format("2006-01-02_15-04-05") + "_" + action.Webhook.RepoName + "_" + after + ".json"

			output := fmt.Sprintf("'%s': Command '%s' executed successfully, output: %s, processing time: %s", action.Webhook.RepoName, command, output, action.ProcessingTime)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           output,
				ProcessingTime: time.Since(startTime),
			})
		}
	}

	action.Webhook.RunCommands(commandHandler)

	action.Success = success
	app.ActionHandler.Add(&action)
}
