package webserver

import (
	"net/http"
	"strconv"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleWebhookRepoNameOperation(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		vars := mux.Vars(r)
		buildName := vars["build_name"]
		operation := vars["operation"]

		webhook, exists := app.WebhookHandler.Webhooks[buildName]
		if !exists {
			http.Error(w, "Repo not found", http.StatusNotFound)
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           "Repo not found",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})
			return
		}

		if operation == "toggle_active" {
			var state bool
			if webhook.Active {
				state = false
			} else {
				state = true
			}

			err := app.WebhookHandler.UpdateActive(buildName, state)
			if err != nil {
				app.Logger.Entry(logger.Container{
					Status:         logger.STATUS_ERROR,
					Error:          "Could not change webhook Active state: " + err.Error(),
					HttpRequest:    r,
					ProcessingTime: time.Since(startTime),
				})
			}

			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           "Webhook state changed to: " + strconv.FormatBool(webhook.Active),
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			return
		}
	}
}
