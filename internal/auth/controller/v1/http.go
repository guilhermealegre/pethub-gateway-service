package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	httpV1 "github.com/guilhermealegre/pethub-gateway-service/api/v1/http"
	"github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
	v1 "github.com/guilhermealegre/pethub-gateway-service/internal/request/domain/v1"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) v1.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()

	httpV1.GetTokenInternalProviders.SetRoute(engine, c.Redirect)
	httpV1.LoginByExternalProvider.SetRoute(engine, c.Redirect)
	httpV1.GetTokenByCallBackExternalProviders.SetRoute(engine, c.Redirect)
	httpV1.SignupInternalProviders.SetRoute(engine, c.Redirect)
	httpV1.SignupInternalProvidersConfirmation.SetRoute(engine, c.Redirect)
	httpV1.CreatePassword.SetRoute(engine, c.Redirect)
	httpV1.Logout.SetRoute(engine, c.Redirect)
	httpV1.Refresh.SetRoute(engine, c.Redirect)

}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	response, body := c.model.Redirect(ctx, config.ServiceEndpoints.AuthEndpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}
