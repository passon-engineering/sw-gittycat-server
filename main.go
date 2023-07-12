package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sw-gittycat-server/modules/application"
	"sw-gittycat-server/modules/git"
	"sw-gittycat-server/modules/webhooks"
	"sw-gittycat-server/modules/webserver"
)

var (
	webhookDefinitionFilePath string
	port                      int
)

func init() {
	flag.StringVar(&webhookDefinitionFilePath, "webhooks", "webhooks.yaml", "Path to the config file")
	flag.IntVar(&port, "port", 3001, "Port to run the server on")
	flag.Parse()
}

func main() {
	app := application.Init()
	webhooks := webhooks.LoadWebhooks(webhookDefinitionFilePath, app)

	for _, webhook := range webhooks {
		git.CloneRepo(webhook.RepoURL, webhook.RepoName, app)
		http.HandleFunc(webhook.Route, webserver.GetWebhookHandler(&webhook))
	}
	git.DeleteAllRepositories(app)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
