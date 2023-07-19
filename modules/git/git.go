package git

import (
	"os/exec"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/file"
	"github.com/passon-engineering/sw-go-utility-lib/system"
)

const (
	GitExitStatusExists = "exit status 128"
)

func CloneRepo(repoUrl string, repoName string, app *application.Application) {
	startTime := time.Now()

	destinationPath := app.ServerPath + "/repositories/" + repoName
	cloneCmd := "git clone " + repoUrl + " " + destinationPath
	output, err := exec.Command("/bin/sh", "-c", cloneCmd).CombinedOutput()

	logEntry := logger.Container{
		Status:         logger.STATUS_INFO,
		ProcessingTime: time.Since(startTime),
	}

	if err != nil {
		logEntry.Info = "Could not clone repository: " + repoUrl

		switch err.Error() {
		case GitExitStatusExists:
			logEntry.Info = "Could not clone repository (already existing): " + repoUrl
		default:
			logEntry.Status = logger.STATUS_ERROR
			logEntry.Info += " " + err.Error() + " " + string(output)
		}
	} else {
		logEntry.Info = "Cloned repository: " + repoUrl + " " + string(output)
	}

	app.Logger.Entry(logEntry)
}

func PullRepo(repoName string, app *application.Application) {
	startTime := time.Now()

	destinationPath := app.ServerPath + "/repositories/" + repoName
	pullCmd := "cd " + destinationPath + " && git pull"

	output, err := system.RunCommand(pullCmd)

	if err != nil {
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_ERROR,
			Error:          "Could not pull repository " + string(output) + " " + err.Error(),
			ProcessingTime: time.Since(startTime),
		})
		return
	}
	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Pulled repository: " + repoName,
		ProcessingTime: time.Since(startTime),
	})
}

func DeleteAllRepositories(app *application.Application) {
	startTime := time.Now()

	ignoreFiles := map[string]bool{
		".gitignore": true,
		// add more files to ignore here
	}

	err := file.DeleteAllExceptIgnored(app.ServerPath+"/repositories/", ignoreFiles)
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_ERROR,
			Error:          "Could not delete all repositories " + err.Error(),
			ProcessingTime: time.Since(startTime),
		})
	}

	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "All repositories deleted",
		ProcessingTime: time.Since(startTime),
	})
}
