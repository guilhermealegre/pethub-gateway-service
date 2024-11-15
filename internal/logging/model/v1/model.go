package v1

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	domainLogging "github.com/guilhermealegre/pethub-gateway-service/internal/logging/domain/v1"
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
