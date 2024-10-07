package v1

import (
	v1Routes "bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	alive "bitbucket.org/asadventure/be-gateway-service/internal/alive/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*domain.DefaultController
	model alive.IModel
}

func NewController(app domain.IApp, model alive.IModel) domain.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	v1Routes.GatewayAlive.SetRoute(c.App().Http().Router(), c.Get)
	v1Routes.GatewayPublicAlive.SetRoute(c.App().Http().Router(), c.GetPublic)
}

/*
	 swagger:route GET /alive alive alive

	 Internal service status check.

		Produces:
		- application/json

		Responses:
		  200: SwaggerAliveResponse
		  400: ErrorResponse
*/
func (c *Controller) Get(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	alive, err := c.model.Get(ctx)
	c.Json(ctx, alive.FromDomainToApi(), err)
}

/*
	 swagger:route GET /p/alive/gateway alive public_alive

	 Public service status check.

		Produces:
		- application/json

		Security:
		  BasicAuth:

		Responses:
		  200: SwaggerPublicAliveResponse
		  400: ErrorResponse
*/
func (c *Controller) GetPublic(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	alive, err := c.model.GetPublic(ctx)
	c.Json(ctx, alive.FromDomainToApi(), err)
}
