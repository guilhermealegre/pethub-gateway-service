package config

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/config"
)

const (
	configFile = "http.yaml"
)

type Endpoint struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type serviceEndpoints struct {
	UserEndpoint     *Endpoint `yaml:"userEndpoint"`
	StoreEndpoint    *Endpoint `yaml:"storeEndpoint"`
	CustomerEndpoint *Endpoint `yaml:"customerEndpoint"`
	OrderEndpoint    *Endpoint `yaml:"orderEndpoint"`
	UploaderEndpoint *Endpoint `yaml:"uploaderEndpoint"`
	LoggingEndpoint  *Endpoint `yaml:"loggingEndpoint"`
}

var ServiceEndpoints serviceEndpoints

func init() {
	_ = config.Load(configFile, &ServiceEndpoints)
}

func GetEndpoint(service string) *Endpoint {
	switch service {
	case "auth":
		return ServiceEndpoints.UserEndpoint
	case "store":
		return ServiceEndpoints.StoreEndpoint
	case "customer":
		return ServiceEndpoints.CustomerEndpoint
	case "order":
		return ServiceEndpoints.OrderEndpoint
	case "uploader":
		return ServiceEndpoints.UploaderEndpoint
	case "logging":
		return ServiceEndpoints.LoggingEndpoint
	}

	return nil
}
