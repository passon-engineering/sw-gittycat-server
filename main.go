package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	yaml "gopkg.in/yaml.v3"
)

type Webhook struct {
	RepoURL   string   `yaml:"repo_url"`
	ClonePath string   `yaml:"clone_path"`
	Webhook   string   `yaml:"webhook"`
	Commands  []string `yaml:"commands"`
}

type Config struct {
	Webhooks []Webhook `yaml:"webhooks"`
}

var (
	configPath string
	port       int
)

func init() {
	flag.StringVar(&configPath, "config", "webhooks.yaml", "Path to the config file")
	flag.IntVar(&port, "port", 3001, "Port to run the server on")
	flag.Parse()
}

func loadConfig() (*Config, error) {
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func cloneRepo(webhook *Webhook) error {
	cloneCmd := "git clone " + webhook.RepoURL + " " + webhook.ClonePath
	fmt.Println(cloneCmd)
	_, err := exec.Command("/bin/sh", "-c", cloneCmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func getWebhookHandler(webhook *Webhook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Webhook received: ", webhook.Webhook)
			for _, command := range webhook.Commands {
				fullCommand := "cd " + webhook.ClonePath + " && " + command
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

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	for _, webhook := range config.Webhooks {
		err := cloneRepo(&webhook)
		if err != nil {
			fmt.Println("Could not clone repo: " + webhook.RepoURL + " " + err.Error())
		}

		http.HandleFunc(webhook.Webhook, getWebhookHandler(&webhook))
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
