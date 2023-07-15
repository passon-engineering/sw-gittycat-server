package webhooks

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

type Webhook struct {
	RepoURL        string   `yaml:"repo_url" json:"repo_url"`
	RepoName       string   `yaml:"repo_name" json:"repo_name"`
	Route          string   `yaml:"route" json:"route"`
	Commands       []string `yaml:"commands" json:"commands"`
	LastCall       string   `yaml:"last_call" json:"last_call"`
	ProcessingTime string   `yaml:"processing_time" json:"processing_time"`
}

type WebhookHandler struct {
	Directory string
	Webhooks  map[string]*Webhook
	sync.Mutex
}

func NewWebhookHandler(directory string) (*WebhookHandler, error) {
	handler := &WebhookHandler{Directory: directory, Webhooks: make(map[string]*Webhook)}
	err := handler.Refresh()
	if err != nil {
		return nil, err
	}

	return handler, nil
}

func (handler *WebhookHandler) Refresh() error {
	handler.Lock()
	defer handler.Unlock()

	files, err := ioutil.ReadDir(handler.Directory)
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
			handler.Webhooks[webhook.Route] = webhook
		}
	}

	return nil
}

func loadWebhook(filePath string) (*Webhook, error) {
	bytes, err := ioutil.ReadFile(filePath)
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

func (handler *WebhookHandler) UpdateLastCall(route string, lastCall string) error {
	handler.Lock()
	defer handler.Unlock()

	webhook, ok := handler.Webhooks[route]
	if !ok {
		return errors.New("webhook not found")
	}

	webhook.LastCall = lastCall
	filePath := filepath.Join(handler.Directory, route+".yaml")
	return saveWebhook(filePath, webhook)
}

func (handler *WebhookHandler) UpdateProcessingTime(route string, processingTime string) error {
	handler.Lock()
	defer handler.Unlock()

	webhook, ok := handler.Webhooks[route]
	if !ok {
		return errors.New("webhook not found")
	}

	webhook.ProcessingTime = processingTime
	filePath := filepath.Join(handler.Directory, route+".yaml")
	return saveWebhook(filePath, webhook)
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

func (webhook *Webhook) RunCommands(handleCommand func(command string, output string, err error)) {
	for _, command := range webhook.Commands {
		output, err := runCommand(command)
		handleCommand(command, output, err)
	}
}

func saveWebhook(filePath string, webhook *Webhook) error {
	data, err := yaml.Marshal(webhook)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

func runCommand(cmdLine string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdLine)
	} else {
		cmd = exec.Command("/bin/sh", "-c", cmdLine)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %s, error: %s, stderr: %s", cmdLine, err, stderr.String())
	} else {
		return stdout.String(), nil
	}
}
