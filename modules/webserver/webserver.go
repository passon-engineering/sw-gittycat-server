package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/webhooks"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/web"
)

func Init(app *application.Application) {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(root(app))

	server := http.Server{
		Handler:      router,
		Addr:         ":" + app.Config.HttpPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	router.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		w.Header().Set("Content-Type", "application/json")
		app.WebhookHandler.Refresh()
		json.NewEncoder(w).Encode(app.WebhookHandler.Webhooks)
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Refreshed and listed available webhooks",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}).Methods("GET")

	router.HandleFunc("/webhooks/{repo_name}/{action}", handleWebhookAction(app))

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

func handleWebhookAction(app *application.Application) func(http.ResponseWriter, *http.Request) {
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
		app.Queue <- webhooks.WebhookAction{Webhook: webhook, Action: action, RequestBody: requestBody}

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Webhook: ",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
