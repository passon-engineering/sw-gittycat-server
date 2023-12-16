package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sw-gittycat-server/modules/webhooks"
	"sync"
	"time"

	"github.com/passon-engineering/sw-go-utility-lib/file"
)

type Action struct {
	Webhook        *webhooks.Webhook      `yaml:"webhook" json:"webhook"`
	Success        bool                   `yaml:"success" json:"success"`
	RequestBody    map[string]interface{} `yaml:"request_body" json:"request_body"`
	LastCall       string                 `yaml:"last_call" json:"last_call"`
	ProcessingTime string                 `yaml:"processing_time" json:"processing_time"`
	FileName       string                 `yaml:"file_name" json:"file_name"`
	Output         string                 `yaml:"output" json:"output"`
}

type ActionHandler struct {
	Directory string
	Actions   map[string]*Action
	sync.Mutex
}

func NewActionHandler(directory string) (*ActionHandler, error) {
	handler := &ActionHandler{Directory: directory, Actions: make(map[string]*Action)}
	err := handler.Reload()
	if err != nil {
		return nil, err
	}

	return handler, nil
}

func (handler *ActionHandler) Reload() error {
	handler.Lock()
	defer handler.Unlock()

	files, err := os.ReadDir(handler.Directory)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	// Create a slice of actions for sorting
	var actionsSlice []*Action
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			filePath := filepath.Join(handler.Directory, file.Name())
			action, err := loadAction(filePath)
			if err != nil {
				return fmt.Errorf("error loading action from %s: %w", file.Name(), err)
			}
			actionsSlice = append(actionsSlice, action)
		}
	}

	// Sort the actions by last call
	sortActionsByLastCall(actionsSlice)

	// Rebuild the map in sorted order
	sortedActions := make(map[string]*Action)
	for _, action := range actionsSlice {
		sortedActions[action.FileName] = action
	}

	handler.Actions = sortedActions

	return nil
}

func loadAction(filePath string) (*Action, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading action file: %w", err)
	}

	var action Action
	err = json.Unmarshal(bytes, &action)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling action: %w", err)
	}

	return &action, nil
}

func sortActionsByLastCall(actions []*Action) {
	sort.Slice(actions, func(i, j int) bool {
		t1, err1 := time.Parse(time.RFC3339, actions[i].LastCall)
		t2, err2 := time.Parse(time.RFC3339, actions[j].LastCall)

		if err1 != nil || err2 != nil {
			return false // or handle the error as appropriate
		}

		return t1.After(t2)
	})
}

func (handler *ActionHandler) Add(a *Action) error {
	handler.Lock()
	defer handler.Unlock()

	handler.Actions[a.FileName] = a

	// Prepare JSON data
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	filePath := filepath.Join(handler.Directory, a.FileName)

	// Create the file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (handler *ActionHandler) DeleteAll() error {
	handler.Lock()
	defer handler.Unlock()

	ignoreFiles := map[string]bool{
		".gitignore": true,
		// add more files to ignore here
	}

	err := file.DeleteAllExceptIgnored(handler.Directory, ignoreFiles)
	if err != nil {
		return err
	} else {
		// Clear the in-memory map
		handler.Actions = make(map[string]*Action)
		return nil
	}
}

func (action *Action) AppendOutput(stringsToAppend ...string) {
	// Concatenate the strings passed as parameters
	appendedString := strings.Join(stringsToAppend, " ")

	action.Output += "\n" + appendedString
}
