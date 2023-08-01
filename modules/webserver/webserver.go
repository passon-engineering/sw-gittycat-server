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
	router := mux.NewRouter()

	server := http.Server{
		Handler:      router,
		Addr:         ":" + app.Config.HttpPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	router.NotFoundHandler = http.HandlerFunc(handleRoot(app))
	router.HandleFunc("/webhooks", handleWebhooks(app))
	router.HandleFunc("/webhooks/refresh", handleWebhooksRefresh(app))
	router.HandleFunc("/webhooks/{repo_name}/{operation}", handleWebhooksRepoNameOperation(app))
	router.HandleFunc("/repositories/stats", handleRepositoriesStats(app))
	router.HandleFunc("/actions", handleActions(app))
	router.HandleFunc("/actions/stats", handleActionsStats(app))

	err := server.ListenAndServe()
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not initialize http server: " + err.Error(),
		})
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}

}
