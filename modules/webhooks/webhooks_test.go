package webhooks

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestWebhookHandler(t *testing.T) {
	// Create a temporary directory for the tests
	tempDir, err := ioutil.TempDir("", "webhookHandlerTest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a new WebhookHandler
	handler, err := NewWebhookHandler(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// Test Add
	webhook := &Webhook{
		RepoURL:        "https://github.com/example/repo.git",
		RepoName:       "example",
		Route:          "webhook1",
		Commands:       []string{"echo hello", "echo world"},
		LastCall:       "",
		ProcessingTime: "",
	}
	err = handler.Add(webhook)
	if err != nil {
		t.Fatal(err)
	}

	// Test Refresh
	err = handler.Refresh()
	if err != nil {
		t.Fatal(err)
	}
	if len(handler.Webhooks) != 1 {
		t.Errorf("Expected 1 webhook, got %d", len(handler.Webhooks))
	}
	if !reflect.DeepEqual(handler.Webhooks[webhook.Route], webhook) {
		t.Errorf("Webhook mismatch: got %+v, want %+v", handler.Webhooks[webhook.Route], webhook)
	}

	// Test UpdateLastCall
	lastCall := time.Now().Format(time.RFC3339)
	err = handler.UpdateLastCall(webhook.Route, lastCall)
	if err != nil {
		t.Fatal(err)
	}
	if handler.Webhooks[webhook.Route].LastCall != lastCall {
		t.Errorf("LastCall mismatch: got %s, want %s", handler.Webhooks[webhook.Route].LastCall, lastCall)
	}

	// Test UpdateProcessingTime
	processingTime := "500ms"
	err = handler.UpdateProcessingTime(webhook.Route, processingTime)
	if err != nil {
		t.Fatal(err)
	}
	if handler.Webhooks[webhook.Route].ProcessingTime != processingTime {
		t.Errorf("ProcessingTime mismatch: got %s, want %s", handler.Webhooks[webhook.Route].ProcessingTime, processingTime)
	}

	// Test Delete
	err = handler.Delete(webhook.Route)
	if err != nil {
		t.Fatal(err)
	}
	if len(handler.Webhooks) != 0 {
		t.Errorf("Expected 0 webhooks, got %d", len(handler.Webhooks))
	}
}
