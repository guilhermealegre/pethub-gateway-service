package v1

import (
	"bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	"bitbucket.org/asadventure/be-gateway-service/api/v1/http/envelope/response"
	"bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	v1 "bitbucket.org/asadventure/be-gateway-service/internal/request/domain/v1"
	user "bitbucket.org/asadventure/be-gateway-service/internal/user/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain/auth"
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
	http.UserScannerList.SetRoute(engine, c.Redirect)
	http.UserScannerLogin.SetRoute(engine, c.Redirect)
	http.UserScannerAuthorize.SetRoute(engine, c.Redirect)
	http.UserLogOff.SetRoute(engine, c.LogOff)
	http.UserByID.SetRoute(engine, c.Redirect)
}

func (c *Controller) Redirect(gCtx *gin.Context) {
	ctx := context.NewContext(gCtx)
	response, body := c.model.Redirect(ctx, config.ServiceEndpoints.UserEndpoint)
	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), body)
}

/*
	 swagger:route PUT /auth/logoff logoff logoff

	 Logoff a session

		Produces:
		- application/json

		Security:
			Bearer:

		Responses:
		  200: SwaggerSuccessResponse
		  400: ErrorResponse
*/
func (c *Controller) LogOff(gCtx *gin.Context) {
	auth.LogOff(gCtx)
	ctx := context.NewContext(gCtx)
	c.Json(ctx, response.SuccessResponse{Success: true}, nil)
}
