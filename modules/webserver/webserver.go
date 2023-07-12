package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"sw-gittycat-server/modules/webhooks"
)

func GetWebhookHandler(webhook webhooks.Webhook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		for _, cmd := range webhook.Commands {
			runCommand(cmd)
		}

		fmt.Fprintf(w, "Webhook %s executed successfully!", webhook.RepoName)
	}
}

func runCommand(cmdLine string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdLine)
	} else {
		cmd = exec.Command("/bin/sh", "-c", cmdLine)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Command failed: %s", cmdLine)
		log.Printf("Error: %s", err)
	} else {
		log.Printf("Command executed successfully: %s", cmdLine)
	}
}
