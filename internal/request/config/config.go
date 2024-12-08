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
	AuthEndpoint         *Endpoint `yaml:"authEndpoint"`
	UserEndpoint         *Endpoint `yaml:"userEndpoint"`
	BookingEndpoint      *Endpoint `yaml:"bookingEndpoint"`
	NotificationEndpoint *Endpoint `yaml:"notificationEndpoint"`
	UploaderEndpoint     *Endpoint `yaml:"uploaderEndpoint"`
	LoggingEndpoint      *Endpoint `yaml:"loggingEndpoint"`
}

var ServiceEndpoints serviceEndpoints

func init() {
	_ = config.Load(configFile, &ServiceEndpoints)
}
