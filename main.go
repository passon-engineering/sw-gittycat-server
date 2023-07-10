package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sw-gittycat-server/application"
	"sw-gittycat-server/modules/git"
	"sw-gittycat-server/modules/webhooks"
	"sw-gittycat-server/modules/webserver"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

var (
	configPath string
	port       int
)

func init() {
	flag.StringVar(&configPath, "webhooks", "webhooks.yaml", "Path to the config file")
	flag.IntVar(&port, "port", 3001, "Port to run the server on")
	flag.Parse()
}

func main() {
	app, err := application.Init()

	webhooks, err := webhooks.LoadWebhooks(configPath)
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not load webhooks: " + err.Error(),
		})
		log.Fatalf("Could not load webhooks: %v", err)
	}

	for _, webhook := range webhooks {
		err := git.CloneRepo(webhook.RepoURL, webhook.ClonePath)
		if err != nil {
			fmt.Println("Could not clone repo: " + webhook.RepoURL + " " + err.Error())
		}

		http.HandleFunc(webhook.Route, webserver.GetWebhookHandler(&webhook))
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
