package webserver

import (
	"net/http"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleNotFound(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Set the header to 404 and write a custom not found message
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))

		// Log the not found event
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Source:         "handleNotFound",
			Info:           "not found",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
