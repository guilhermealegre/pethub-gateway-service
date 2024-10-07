package v1

import (
	"context"

	domainLogging "bitbucket.org/asadventure/be-gateway-service/internal/logging/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	logging "bitbucket.org/asadventure/be-logging-service/api/v1/grpc/logging_service_logging"
)

type Streaming struct {
	app           domain.IApp
	loggingClient logging.LoggingClient
}

func NewStreaming(app domain.IApp, loggingClient logging.LoggingClient) domainLogging.IStreaming {
	return &Streaming{
		app:           app,
		loggingClient: loggingClient,
	}
}

func (s *Streaming) Log(message []byte) (err error) {
	_, err = s.loggingClient.Log(context.Background(), &logging.LogRequest{Message: message})
	return err
}
