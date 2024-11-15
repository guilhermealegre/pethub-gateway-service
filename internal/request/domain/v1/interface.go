package v1

import (
	"net/http"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	"github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
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
