package actions

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sw-gittycat-server/modules/webhooks"
)

type Action struct {
	Webhook        *webhooks.Webhook      `yaml:"webhook" json:"webhook"`
	Success        bool                   `yaml:"success" json:"success"`
	RequestBody    map[string]interface{} `yaml:"request_body" json:"request_body"`
	LastCall       string                 `yaml:"last_call" json:"last_call"`
	ProcessingTime string                 `yaml:"processing_time" json:"processing_time"`
}

func (a *Action) Create(serverPath string) error {
	after, ok := a.RequestBody["after"].(string)
	if !ok || len(after) == 0 {
		return errors.New("invalid or missing 'after' field in request body")
	}

	// Prepare JSON data
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	// Create the file name
	fileName := a.Webhook.RepoName + "_SHA_" + after + ".json"
	filePath := filepath.Join(serverPath, "actions", fileName)

	// Create the file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
