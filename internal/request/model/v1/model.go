package v1

import (
	"encoding/json"
	"fmt"
	dCtx "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain/context"
	"io"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	"github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
	requestDomain "github.com/guilhermealegre/pethub-gateway-service/internal/request/domain/v1"
)

type Model struct {
	app    domain.IApp
	client requestDomain.HttpClient
}

func NewModel(app domain.IApp, client requestDomain.HttpClient) requestDomain.IModel {
	return &Model{
		app:    app,
		client: client,
	}
}

func (m *Model) Redirect(ctx dCtx.IContext, serviceEndpoint *config.Endpoint) (*http.Response, []byte) {
	var err error
	ctx.Request().URL, err = m.prepareRedirectUrl(ctx.Request(), serviceEndpoint)
	if err != nil {
		response := &http.Response{StatusCode: http.StatusInternalServerError}
		rBytes, errMarshal := json.Marshal(err)
		if errMarshal != nil {
			m.app.Logger().Log().Do(errMarshal)
		}
		return response, rBytes
	}
	ctx.Request().RequestURI = ""

	response, body := m.getResponse(ctx.Request())

	// Set Cookies from Redirect
	for _, cookie := range response.Cookies() {
		cookieRedirect := &http.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Expires:  cookie.Expires,
			Path:     cookie.Path,
			SameSite: cookie.SameSite,
			HttpOnly: cookie.HttpOnly,
			Secure:   cookie.Secure,
		}
		// Set the cookie to the response
		http.SetCookie(ctx.Response(), cookieRedirect)
	}

	return response, body

}

func (m *Model) prepareRedirectUrl(request *http.Request, serviceEndpoint *config.Endpoint) (*url.URL, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s:%s%s",
		serviceEndpoint.Protocol, serviceEndpoint.Host, serviceEndpoint.Port, request.URL.Path))
	if err != nil {
		m.app.Logger().Log().Do(err)
		return nil, err
	}

	u.RawQuery = request.URL.RawQuery

	return u, nil
}

func (m *Model) getResponse(request *http.Request) (response *http.Response, rBytes []byte) {
	// due o http client check content length after body close, we must send the body with close allowance
	//request.Body = httpInfra.NewReader(request.Body, true) TODO: review this code

	// sending the span information in order to ensure continuity of the same trace
	if span := trace.SpanFromContext(request.Context()); span != nil {
		otel.GetTextMapPropagator().Inject(request.Context(), propagation.HeaderCarrier(request.Header))
	}

	response, err := m.client.Do(request)
	if err != nil {
		m.app.Logger().Log().Do(err)
		response = &http.Response{StatusCode: http.StatusInternalServerError}
		var errMarshal error
		rBytes, errMarshal = json.Marshal(err)
		if errMarshal != nil {
			m.app.Logger().Log().Do(errMarshal)
		}
		return response, rBytes
	}

	rBytes, err = io.ReadAll(response.Body)
	if err != nil {
		m.app.Logger().Log().Do(err)
		response = &http.Response{StatusCode: http.StatusInternalServerError}
		var errMarshal error
		rBytes, errMarshal = json.Marshal(err)
		if errMarshal != nil {
			m.app.Logger().Log().Do(errMarshal)
		}
		return response, rBytes
	}

	return response, rBytes
}
