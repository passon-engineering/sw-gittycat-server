package actions

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (a *Action) Save() error {
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
	fileName := filepath.Join("/actions", fmt.Sprintf("%s.json", after))

	// Create the file
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
