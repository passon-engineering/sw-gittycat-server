package webserver

import (
	"net/http"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleWebhooksRefresh(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		app.WebhookHandler.Reload()

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Webhooks refreshed",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
