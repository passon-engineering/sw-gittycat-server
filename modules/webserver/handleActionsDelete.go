package webserver

import (
	"net/http"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleActionsDelete(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		err := app.ActionHandler.DeleteAll()
		if err != nil {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Info:           "Failed to delete available actions",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			w.WriteHeader(http.StatusInternalServerError)
		} else {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           "Deleted available actions",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			w.WriteHeader(http.StatusOK)
		}
	}
}
