package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/webhooks"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/web"
)

func Init(app *application.Application) {
	router := mux.NewRouter()

	server := http.Server{
		Handler:      router,
		Addr:         ":" + app.Config.HttpPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	router.NotFoundHandler = http.HandlerFunc(root(app))
	router.HandleFunc("/webhooks", handleWebhooks(app))
	router.HandleFunc("/webhooks/refresh", handleWebhooksReload(app))
	router.HandleFunc("/webhooks/{repo_name}/{action}", handleWebhooksRepoNameAction(app))

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

func root(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		path := r.URL.Path[1:]
		if path == "" {
			http.Redirect(w, r, "/index.html", http.StatusSeeOther)

			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Source:         "root_redirect",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})
			return
		}

		web.ServeStaticFile(w, r, "frontend/dist/"+path)

	}
}

func handleWebhooks(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app.WebhookHandler.Webhooks)
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Refreshed and listed available webhooks",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}

func handleWebhooksReload(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		app.WebhookHandler.Reload()

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Webhooks reloaded",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}

func handleWebhooksRepoNameAction(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		vars := mux.Vars(r)
		repoName := vars["repo_name"]
		action := vars["action"]

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

		if action == "toggle_active" {
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
		app.Queue <- webhooks.WebhookAction{Webhook: webhook, RequestBody: requestBody}

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Webhook: ",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
