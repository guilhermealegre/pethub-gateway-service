package v1

import (
	domainLogging "bitbucket.org/asadventure/be-gateway-service/internal/logging/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
)

type Model struct {
	app       domain.IApp
	streaming domainLogging.IStreaming
}

func NewModel(app domain.IApp, streaming domainLogging.IStreaming) domainLogging.IModel {
	return &Model{
		app:       app,
		streaming: streaming,
	}
}

func (m *Model) Log(message []byte) error {
	return m.streaming.Log(message)
}
