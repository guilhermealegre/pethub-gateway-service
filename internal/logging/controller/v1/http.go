package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	"github.com/guilhermealegre/pethub-gateway-service/api/v1/http"
	user "github.com/guilhermealegre/pethub-gateway-service/internal/logging/domain/v1"
	"github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
	v1 "github.com/guilhermealegre/pethub-gateway-service/internal/request/domain/v1"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) user.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()
	http.LoggingCreateFeLog.SetRoute(engine, c.Redirect)
}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	response, body := c.model.Redirect(ctx, config.ServiceEndpoints.LoggingEndpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}
