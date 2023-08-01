package main

import (
	"flag"
	"strconv"
	"sw-gittycat-server/modules/actions"
	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/git"
	"sw-gittycat-server/modules/processor"
	"sw-gittycat-server/modules/webhooks"
	"sw-gittycat-server/modules/webserver"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

var (
	webhookDefinitionFilePath string
	port                      int
)

func init() {
	flag.StringVar(&webhookDefinitionFilePath, "webhooks", "webhooks.yaml", "Path to the config file")
	flag.IntVar(&port, "port", 3001, "Port to run the server on")
	flag.Parse()
}

func main() {
	app := application.Init()

	var err error

	app.WebhookHandler, err = webhooks.NewWebhookHandler(app.ServerPath + "/webhooks/")
	if err != nil {
		startTime := time.Now()
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_ERROR,
			Error:          "Could not initialize list of webhooks " + err.Error(),
			ProcessingTime: time.Since(startTime),
		})
	} else {
		startTime := time.Now()
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Initialized list of webhooks:",
			ProcessingTime: time.Since(startTime),
		})
		for _, webhook := range app.WebhookHandler.Webhooks {
			app.Logger.Entry(logger.Container{
				Status: logger.STATUS_INFO,
				Info:   webhook.RepoName + " listening to: " + webhook.Route,
			})
		}
	}

	for _, webhook := range app.WebhookHandler.Webhooks {
		git.CloneRepo(webhook.RepoURL, webhook.RepoName, app)
	}

	app.ActionHandler, err = actions.NewActionHandler(app.ServerPath + "/actions")
	if err != nil {
		startTime := time.Now()
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_ERROR,
			Error:          "Could not initialize list of webhooks " + err.Error(),
			ProcessingTime: time.Since(startTime),
		})
	} else {
		startTime := time.Now()
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Initialized list of webhooks:",
			ProcessingTime: time.Since(startTime),
		})
		for _, webhook := range app.WebhookHandler.Webhooks {
			app.Logger.Entry(logger.Container{
				Status: logger.STATUS_INFO,
				Info:   webhook.RepoName + " listening to: " + webhook.Route,
			})
		}
	}

	processor.Init(app)

	app.Config.HttpPort = strconv.Itoa(port)
	app.Config.WebDirectory = "/frontend/dist/"
	webserver.Init(app)

	//git.DeleteAllRepositories(app)

}
