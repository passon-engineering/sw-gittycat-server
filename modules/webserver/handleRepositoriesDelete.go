package webserver

import (
	"net/http"
	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/git"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleRepositoriesDelete(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		err := git.DeleteAllRepositories(app)
		if err != nil {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Info:           "Failed to delete available repositories",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			w.WriteHeader(http.StatusInternalServerError)
		} else {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Info:           "Deleted available repositories",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})

			w.WriteHeader(http.StatusOK)
		}
	}
}
