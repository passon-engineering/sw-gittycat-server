package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sw-gittycat-server/modules/git"
	"sw-gittycat-server/modules/webhooks"
	"sw-gittycat-server/modules/webserver"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/networking"
)

var (
	configPath string
	port       int
)

func init() {
	flag.StringVar(&configPath, "config", "webhooks.yaml", "Path to the config file")
	flag.IntVar(&port, "port", 3001, "Port to run the server on")
	flag.Parse()
}

func main() {
	appLogger, err := logger.NewLogger(
		[]logger.LogFormat{
			logger.FORMAT_TIMESTAMP,
			logger.FORMAT_STATUS,
			logger.FORMAT_PRE_TEXT,
			logger.FORMAT_HTTP_REQUEST,
			logger.FORMAT_ID,
			logger.FORMAT_SOURCE,
			logger.FORMAT_DATA,
			logger.FORMAT_ERROR,
			logger.FORMAT_PROCESSING_TIME,
		}, logger.Options{
			OutputToStdout:   true,
			OutputToFile:     true,
			OutputFolderPath: "/var/log/gittycat/",
		}, logger.Container{
			Status: logger.STATUS_INFO,
			Info:   "System Logger succesfully started! Awaiting logger tasks...",
		})
	if err != nil {
		log.Fatalf(err.Error())
	}

	ip, err := networking.GetNetworkExternalIP()
	if err != nil {
		appLogger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not get network external IP: " + err.Error(),
		})
		log.Fatalf("Could not get network external IP: %v", err)
	}
	appLogger.Entry(logger.Container{
		Status: logger.STATUS_INFO,
		Info:   "Network external IP: " + ip,
	})

	webhooks, err := webhooks.LoadWebhooks(configPath)
	if err != nil {
		appLogger.Entry(logger.Container{
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
