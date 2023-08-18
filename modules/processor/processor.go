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

	for _, repo := range action.Webhook.Repos {
		git.CloneRepo(repo.RepoURL, repo.RepoName, app)
		action.AppendOutput("cloned repository:", repo.RepoName)
		git.PullRepo(repo.RepoName, app)
		action.AppendOutput("pulled repository:", repo.RepoName)
	}

	commandHandlerInner := func(repoName string, command string, output string, err error) {
		if err != nil {
			success = false
			output := fmt.Sprintf("'%s': Command '%s' failed: %v", repoName, command, err)

			action.AppendOutput(output)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Error:          output,
				ProcessingTime: time.Since(startTime),
			})
		} else {
			output := fmt.Sprintf("'%s': Command '%s' executed successfully, output: %s", repoName, command, output)

			action.AppendOutput(output)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           output,
				ProcessingTime: time.Since(startTime),
			})
		}
	}
	action.AppendOutput("INNER COMMANDS")
	action.Webhook.RunInnerRepoCommands(commandHandlerInner)

	commandHandlerComposer := func(command string, output string, err error) {
		if err != nil {
			success = false
			output := fmt.Sprintf("'%s': Command '%s' failed: %v", action.Webhook.BuildName, command, err)

			action.AppendOutput(output)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Error:          output,
				ProcessingTime: time.Since(startTime),
			})
		} else {
			output := fmt.Sprintf("'%s': Command '%s' executed successfully, output: %s", action.Webhook.BuildName, command, output)

			action.AppendOutput(output)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           output,
				ProcessingTime: time.Since(startTime),
			})
		}
	}
	action.AppendOutput("COMPOSER COMMANDS")
	action.Webhook.RunComposerCommands(commandHandlerComposer)

	action.ProcessingTime = time.Since(startTime).String()
	action.Success = success

	after, ok := action.RequestBody["after"].(string)
	if !ok || len(after) == 0 {
		after = "no_hash_available"
	}

	action.FileName = startTime.Format("2006-01-02_15-04-05") + "_" + action.Webhook.BuildName + "_" + after + ".json"

	app.ActionHandler.Add(&action)
}
