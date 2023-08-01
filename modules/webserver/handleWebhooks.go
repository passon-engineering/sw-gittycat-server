package webserver

import (
	"encoding/json"
	"net/http"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleWebhooks(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app.WebhookHandler.Webhooks)
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Listed available webhooks",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
