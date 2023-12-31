package webserver

import (
	"encoding/json"
	"math"
	"net/http"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/file"
)

func handleArtifactsStats(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		w.Header().Set("Content-Type", "application/json")

		stats, err := file.CountFilesAndFolders(app.ServerPath+"/artifacts", math.MaxInt64)
		if err != nil {
			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_ERROR,
				Error:          "Could not read artifacts stats: " + err.Error(),
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})
		}

		json.NewEncoder(w).Encode(stats)
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Artifacts stats transmitted",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
