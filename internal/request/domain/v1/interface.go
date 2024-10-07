package v1

import (
	"net/http"

	"bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
)

type IController interface {
	domain.IController
}

type IModel interface {
	Redirect(ctx domain.IContext, serviceEndpoint *config.Endpoint) (*http.Response, []byte)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
