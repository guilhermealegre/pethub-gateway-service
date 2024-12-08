package v1

import (
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"net/http"

	"github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
)

type IController interface {
	domain.IController
}

type IModel interface {
	Redirect(ctx dCtx.IContext, serviceEndpoint *config.Endpoint) (*http.Response, []byte)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
