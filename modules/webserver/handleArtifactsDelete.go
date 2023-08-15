package webserver

import (
	"net/http"
	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/artifacts"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleArtifactsDelete(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		err := artifacts.DeleteAll(app)
		if err != nil {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Info:           "Failed to delete available artifacts",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			w.WriteHeader(http.StatusInternalServerError)
		} else {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           "Deleted available artifacts",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			w.WriteHeader(http.StatusOK)
		}
	}
}
