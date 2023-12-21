package webserver

import (
	"net/http"
	"os"
	"time"

	"sw-gittycat-server/modules/application"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func Init(app *application.Application) {
	// Main router for the first server
	webInterfaceRouter := mux.NewRouter()
	webInterfaceRouter.NotFoundHandler = http.HandlerFunc(handleRoot(app))
	webInterfaceRouter.HandleFunc("/webhooks", handleWebhooks(app))
	webInterfaceRouter.HandleFunc("/webhooks/refresh", handleWebhooksRefresh(app))
	webInterfaceRouter.HandleFunc("/webhook/{build_name}/{operation}", handleWebhookRepoNameOperation(app))
	webInterfaceRouter.HandleFunc("/repositories/stats", handleRepositoriesStats(app))
	webInterfaceRouter.HandleFunc("/repositories/delete", handleRepositoriesDelete(app))
	webInterfaceRouter.HandleFunc("/actions/{page}", handleActions(app))
	webInterfaceRouter.HandleFunc("/actions/stats", handleActionsStats(app))
	webInterfaceRouter.HandleFunc("/actions/delete", handleActionsDelete(app))
	webInterfaceRouter.HandleFunc("/artifacts/delete", handleArtifactsDelete(app))
	webInterfaceRouter.HandleFunc("/artifacts/stats", handleArtifactsStats(app))

	// Create the first server
	mainServer := http.Server{
		Handler:      webInterfaceRouter,
		Addr:         ":" + app.Config.WebInterfaceHttpPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go func() {
		if err := mainServer.ListenAndServe(); err != nil {
			app.Logger.Entry(logger.Container{
				Status: logger.STATUS_ERROR,
				Error:  "Could not initialize http server: " + err.Error(),
			})
			time.Sleep(2 * time.Second)
			os.Exit(1)
		}
	}()

	// Router for the webhook
	webhookRouter := mux.NewRouter()
	webhookRouter.HandleFunc("/webhook/{build_name}/{operation}", handleWebhook(app))

	// Create the webhook server
	webhookServer := http.Server{
		Handler:      webhookRouter,
		Addr:         ":" + app.Config.WebhookHttpPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go func() {
		if err := webhookServer.ListenAndServe(); err != nil {
			app.Logger.Entry(logger.Container{
				Status: logger.STATUS_ERROR,
				Error:  "Could not initialize http server: " + err.Error(),
			})
			time.Sleep(2 * time.Second)
			os.Exit(1)
		}
	}()
}
