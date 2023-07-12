package webhooks

import (
	"io/ioutil"
	"sw-gittycat-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
	"gopkg.in/yaml.v3"
)

type Webhook struct {
	RepoURL  string   `yaml:"repo_url"`
	RepoName string   `yaml:"repo_name"`
	Route    string   `yaml:"route"`
	Commands []string `yaml:"commands"`
}

type FileContent struct {
	Webhooks []Webhook `yaml:"webhooks"`
}

func LoadWebhooks(webhookDefinitionFilePath string, app *application.Application) []Webhook {
	startTime := time.Now()

	filePath := app.ServerPath + "/modules/" + webhookDefinitionFilePath

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_ERROR,
			Error:          "Could not read list of webhooks " + err.Error(),
			ProcessingTime: time.Since(startTime),
		})
		return nil
	}

	var fileContent FileContent
	err = yaml.Unmarshal(bytes, &fileContent)
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_ERROR,
			Error:          "Could not parse list of webhooks " + err.Error(),
			ProcessingTime: time.Since(startTime),
		})
		return nil
	}

	app.Logger.Entry(logger.Container{
		Status:         logger.STATUS_INFO,
		Info:           "Initialized list of webhooks",
		ProcessingTime: time.Since(startTime),
	})

	return fileContent.Webhooks
}
