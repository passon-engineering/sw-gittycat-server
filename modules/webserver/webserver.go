package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"sw-gittycat-server/modules/webhooks"
)

func GetWebhookHandler(webhook *webhooks.Webhook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Webhook received: ", webhook.Route)
			for _, command := range webhook.Commands {
				fullCommand := "cd " + webhook.RepoName + " && " + command
				fmt.Println(fullCommand)
				out, err := exec.Command("/bin/sh", "-c", fullCommand).Output()
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Printf("%s", out)
			}
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
