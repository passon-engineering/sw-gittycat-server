package webserver

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"sw-gittycat-server/modules/application"

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

	router.HandleFunc("/all_available_webhooks", func(w http.ResponseWriter, r *http.Request) {
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
	})

	router.HandleFunc("/webhooks/{key}", func(w http.ResponseWriter, r *http.Request) {

	})

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
		start := time.Now()

		path := r.URL.Path[1:]
		if path == "" {
			http.Redirect(w, r, "/index.html", http.StatusSeeOther)

			app.Logger.Entry(logger.Container{
				Source:         "root_redirect",
				HttpRequest:    r,
				ProcessingTime: time.Since(start),
			})
			return
		}

		web.ServeStaticFile(w, r, "frontend/dist/"+path)

	}
}
