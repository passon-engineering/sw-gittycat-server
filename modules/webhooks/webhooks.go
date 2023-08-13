package webhooks

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/passon-engineering/sw-go-utility-lib/system"
	"gopkg.in/yaml.v3"
)

type Webhook struct {
	BuildName        string   `yaml:"build_name"`
	Active           bool     `yaml:"active"`
	Route            string   `yaml:"route"`
	ComposerCommands []string `yaml:"composer_commands"`
	Repos            []Repo   `yaml:"repos"`
}

type Repo struct {
	RepoURL            string   `yaml:"repo_url"`
	RepoName           string   `yaml:"repo_name"`
	BranchOrCommitHash string   `yaml:"branch_or_commit_hash"`
	InnerRepoCommands  []string `yaml:"inner_repo_commands"`
}

type WebhookHandler struct {
	Directory string
	Webhooks  map[string]*Webhook
	sync.Mutex
}

func NewWebhookHandler(directory string) (*WebhookHandler, error) {
	handler := &WebhookHandler{Directory: directory, Webhooks: make(map[string]*Webhook)}
	err := handler.Reload()
	if err != nil {
		return nil, err
	}

	return handler, nil
}

func (handler *WebhookHandler) Reload() error {
	handler.Lock()
	defer handler.Unlock()

	files, err := os.ReadDir(handler.Directory)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	handler.Webhooks = make(map[string]*Webhook)

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			filePath := filepath.Join(handler.Directory, file.Name())
			webhook, err := loadWebhook(filePath)
			if err != nil {
				return fmt.Errorf("error loading webhook from %s: %w", file.Name(), err)
			}
			handler.Webhooks[webhook.BuildName] = webhook
		}
	}

	return nil
}

func loadWebhook(filePath string) (*Webhook, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading webhook file: %w", err)
	}

	var webhook Webhook
	err = yaml.Unmarshal(bytes, &webhook)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling webhook: %w", err)
	}

	return &webhook, nil
}

func (handler *WebhookHandler) Delete(route string) error {
	handler.Lock()
	defer handler.Unlock()

	_, ok := handler.Webhooks[route]
	if !ok {
		return errors.New("webhook not found")
	}

	filePath := filepath.Join(handler.Directory, route+".yaml")
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	delete(handler.Webhooks, route)

	return nil
}

func (handler *WebhookHandler) Add(webhook *Webhook) error {
	handler.Lock()
	defer handler.Unlock()

	_, ok := handler.Webhooks[webhook.Route]
	if ok {
		return errors.New("webhook already exists")
	}

	filePath := filepath.Join(handler.Directory, webhook.Route+".yaml")
	err := saveWebhook(filePath, webhook)
	if err != nil {
		return err
	}

	handler.Webhooks[webhook.Route] = webhook
	return nil
}

func (handler *WebhookHandler) UpdateActive(route string, status bool) error {
	handler.Lock()
	defer handler.Unlock()

	webhook, ok := handler.Webhooks[route]
	if !ok {
		return errors.New("webhook not found")
	}

	webhook.Active = status
	filePath := filepath.Join(handler.Directory, route+".yaml")
	return saveWebhook(filePath, webhook)
}

func (webhook *Webhook) RunInnerRepoCommands(handleCommand func(repoName string, command string, output string, err error)) {
	for _, repo := range webhook.Repos {
		for _, command := range repo.InnerRepoCommands {
			output, err := system.RunCommand("cd " + repo.RepoName + " && " + command)
			handleCommand(repo.RepoName, command, output, err)
		}
	}
}

func (webhook *Webhook) RunComposerCommands(handleCommand func(command string, output string, err error)) {
	for _, command := range webhook.ComposerCommands {
		output, err := system.RunCommand(command)
		handleCommand(command, output, err)
	}
}

func saveWebhook(filePath string, webhook *Webhook) error {
	data, err := yaml.Marshal(webhook)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
