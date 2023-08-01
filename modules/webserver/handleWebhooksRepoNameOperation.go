package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sw-gittycat-server/modules/actions"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleWebhooksRepoNameOperation(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		vars := mux.Vars(r)
		repoName := vars["repo_name"]
		operation := vars["operation"]

		webhook, exists := app.WebhookHandler.Webhooks[repoName]
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

			err := app.WebhookHandler.UpdateActive(repoName, state)
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

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			// Handle the error
			// You can return an error response or take appropriate action
			fmt.Println("Error reading request body:", err.Error())
			return
		}

		// Parse the request body into a map[string]interface{}
		var requestBody map[string]interface{}
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			// Handle the error
			// You can return an error response or take appropriate action
			fmt.Println("Error parsing request body:", err.Error())
			requestBody = nil
		}

		// Add the webhook action to the queue
		app.Queue <- actions.Action{Webhook: webhook, RequestBody: requestBody}

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Webhook: ",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
