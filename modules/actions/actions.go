package actions

import "sw-gittycat-server/modules/webhooks"

type Action struct {
	Webhook        *webhooks.Webhook      `yaml:"webhook" json:"webhook"`
	Success        bool                   `yaml:"success" json:"success"`
	RequestBody    map[string]interface{} `yaml:"request_body" json:"request_body"`
	LastCall       string                 `yaml:"last_call" json:"last_call"`
	ProcessingTime string                 `yaml:"processing_time" json:"processing_time"`
}
