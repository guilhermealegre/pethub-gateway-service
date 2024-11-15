package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
	v1Auth "github.com/guilhermealegre/pethub-gateway-service/internal/auth/domain/v1"
	"github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
	v1 "github.com/guilhermealegre/pethub-gateway-service/internal/request/domain/v1"
)

type Controller struct {
	*domain.DefaultController
	model v1.IModel
}

func NewController(app domain.IApp, model v1.IModel) v1Auth.IController {
	return &Controller{
		DefaultController: domain.NewDefaultController(app),
		model:             model,
	}
}

func (c *Controller) Register() {
	_ = c.App().Http().Router()

}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	response, body := c.model.Redirect(ctx, config.ServiceEndpoints.UserEndpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}
