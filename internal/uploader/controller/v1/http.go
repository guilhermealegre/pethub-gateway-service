package v1

import (
	"bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	"bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	v1 "bitbucket.org/asadventure/be-gateway-service/internal/request/domain/v1"
	user "bitbucket.org/asadventure/be-gateway-service/internal/uploader/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
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
	http.ImageUpload.SetRoute(engine, c.Redirect)
	http.ImagePUpload.SetRoute(engine, c.Redirect)
}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	response, body := c.model.Redirect(ctx, config.ServiceEndpoints.UploaderEndpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}
