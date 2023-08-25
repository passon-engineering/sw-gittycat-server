package webserver

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"sw-gittycat-server/modules/actions"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleActions(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		const ActionsPerPage = 0 // This defines how many actions should be in one page

		vars := mux.Vars(r)
		pageStr := vars["page"]
		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}

		// Create a slice of keys and sort them
		keys := make([]string, 0, len(app.ActionHandler.Actions))
		for key := range app.ActionHandler.Actions {
			keys = append(keys, key)
		}
		sort.Strings(keys) // Assuming alphanumeric order, adjust as necessary

		startIndex := (page - 1) * ActionsPerPage
		endIndex := page * ActionsPerPage
		if startIndex >= len(keys) {
			http.Error(w, "Page number out of range", http.StatusNotFound)
			return
		}

		if endIndex > len(keys) {
			endIndex = len(keys)
		}

		// Return actions based on sorted keys in the pagination range
		actionsToReturn := make(map[string]*actions.Action)
		for _, key := range keys[startIndex:endIndex] {
			actionsToReturn[key] = app.ActionHandler.Actions[key]
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(actionsToReturn)

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Listed available actions",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
