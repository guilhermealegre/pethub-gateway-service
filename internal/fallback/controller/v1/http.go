package v1

import (
	"bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	customer "bitbucket.org/asadventure/be-gateway-service/internal/customer/domain/v1"
	"bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	v1 "bitbucket.org/asadventure/be-gateway-service/internal/request/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) customer.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()
	http.PublicAlive.SetRoute(engine, c.Redirect)
	http.SwaggerJson.SetRoute(engine, c.Redirect)
}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)

	endpoint := config.GetEndpoint(gCtx.Param("service"))
	if endpoint == nil {
		return
	}

	response, body := c.model.Redirect(ctx, endpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}
