package webserver

import (
	"net/http"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"github.com/passon-engineering/sw-go-utility-lib/web"
)

func handleRoot(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		path := r.URL.Path[1:]
		if path == "" {
			http.Redirect(w, r, "/index.html", http.StatusSeeOther)

			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Source:         "root_redirect",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})
			return
		}

		web.ServeStaticFile(w, r, app.ServerPath+"/frontend/dist/"+path)

	}
}
