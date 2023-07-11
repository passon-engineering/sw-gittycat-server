package webhooks

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Webhook struct {
	RepoURL         string   `yaml:"repo_url"`
	DestinationName string   `yaml:"destination_name"`
	Route           string   `yaml:"route"`
	Commands        []string `yaml:"commands"`
}

type FileContent struct {
	Webhooks []Webhook `yaml:"webhooks"`
}

func LoadWebhooks(filePath string) ([]Webhook, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var fileContent FileContent
	err = yaml.Unmarshal(bytes, &fileContent)
	if err != nil {
		return nil, err
	}

	return fileContent.Webhooks, nil
}
