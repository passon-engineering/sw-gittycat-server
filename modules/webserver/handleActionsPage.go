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

		const ActionsPerPage = 5 // This defines how many actions should be in one page

		vars := mux.Vars(r)
		pageStr := vars["page"]
		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}

		// Create a slice of actions
		actionsSlice := make([]*actions.Action, 0, len(app.ActionHandler.Actions))
		for _, action := range app.ActionHandler.Actions {
			actionsSlice = append(actionsSlice, action)
		}

		// Sort actions by the LastCall date
		sort.Slice(actionsSlice, func(i, j int) bool {
			t1, err1 := time.Parse(time.RFC3339, actionsSlice[i].LastCall)
			t2, err2 := time.Parse(time.RFC3339, actionsSlice[j].LastCall)

			if err1 != nil || err2 != nil {
				// Handle parsing error if necessary
				return false
			}

			return t1.After(t2) // Change to t1.Before(t2) if you want oldest first
		})

		startIndex := (page - 1) * ActionsPerPage
		endIndex := page * ActionsPerPage
		if endIndex > len(actionsSlice) {
			endIndex = len(actionsSlice)
		}

		// Return actions based on sorted slice in the pagination range
		actionsToReturn := make(map[string]*actions.Action)
		for _, action := range actionsSlice[startIndex:endIndex] {
			actionsToReturn[action.FileName] = action // Assuming FileName is the key
		}

		// Calculate total pages
		totalActions := len(actionsSlice)
		totalPages := totalActions / ActionsPerPage
		if totalActions%ActionsPerPage != 0 {
			totalPages++
		}

		// Create a response structure
		response := struct {
			Data       map[string]*actions.Action `json:"data"`
			TotalPages int                        `json:"totalPages"`
		}{
			Data:       actionsToReturn,
			TotalPages: totalPages,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Info:           "Listed available actions",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
