package v1

import (
	v1Routes "bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	access "bitbucket.org/asadventure/be-gateway-service/internal/access/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*domain.DefaultController
	model access.IModel
}

func NewController(app domain.IApp, model access.IModel) domain.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	v1Routes.GatewayPublicAccesssClearance.SetRoute(c.App().Http().Router(), c.Get)
}

/*
	swagger:route GET /p/access-clearance access-clearance access-clearance

	Internal service access clearance check.


	Produces:
	- application/json

	Responses:
	  200: SwaggerAccessClearanceResponse
	  400: ErrorResponse
*/

func (c *Controller) Get(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	accessClearance, err := c.model.Get(ctx)
	c.Json(ctx, accessClearance.FromDomainToApi(), err)
}
