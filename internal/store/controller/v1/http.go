package v1

import (
	"bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	"bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	v1 "bitbucket.org/asadventure/be-gateway-service/internal/request/domain/v1"
	store "bitbucket.org/asadventure/be-gateway-service/internal/store/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) store.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	engine := c.App().Http().Router()
	http.StoreList.SetRoute(engine, c.Redirect)
	http.StoreByIP.SetRoute(engine, c.Redirect)
	http.StoreByID.SetRoute(engine, c.Redirect)
}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	response, body := c.model.Redirect(ctx, config.ServiceEndpoints.StoreEndpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}
