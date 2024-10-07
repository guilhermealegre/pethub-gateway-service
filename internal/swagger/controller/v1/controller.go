package v1

import (
	"fmt"
	"strconv"
	"strings"

	v1Routes "bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gin-gonic/gin"
)

const v1 = "v1"

type Controller struct {
	domain.IController
	app domain.IApp
}

func NewController(app domain.IApp) domain.IController {
	return &Controller{
		app: app,
	}
}

func (c *Controller) Register() {
	v1Int, _ := strconv.Atoi(strings.TrimPrefix(v1, "v"))

	v1Router := c.app.Http().Router().Group("v1")
	v1Router.StaticFile(c.StaticFile(v1Int))
	v1Routes.SwaggerSwagger.SetRoute(c.app.Http().Router(), c.Swagger(v1Int))
	v1Routes.SwaggerDocs.SetRoute(c.app.Http().Router(), c.Docs(v1Int))
}

/*
	 swagger:route GET /:service/docs swagger docs

	 Get swagger docs

		Produces:
		- text/html

		Responses:
		  200:
		  400: ErrorResponse
*/
func (c *Controller) Docs(version int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		service := ctx.Param("service")
		gin.WrapH(middleware.Redoc(
			middleware.RedocOpts{
				Path:    fmt.Sprintf("v%d/p/documentation/%s/docs", version, service),
				SpecURL: fmt.Sprintf("/v%d/p/documentation/%s/swagger.json", version, service),
			}, nil))(ctx)
	}
}

/*
	 swagger:route GET /:service/swagger swagger swagger

	 Get swagger

		Produces:
		- text/html

		Responses:
		  200:
		  400: ErrorResponse
*/
func (c *Controller) Swagger(version int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		service := ctx.Param("service")
		gin.WrapH(middleware.SwaggerUI(
			middleware.SwaggerUIOpts{
				Path:    fmt.Sprintf("v%d/p/documentation/%s/swagger", version, service),
				SpecURL: fmt.Sprintf("/v%d/p/documentation/%s/swagger.json", version, service),
			}, nil))(ctx)
	}
}

func (c *Controller) StaticFile(version int) (relativePath, filePath string) {
	return "/p/documentation/gateway/swagger.json", fmt.Sprintf("internal/swagger/docs/v%d/swagger.json", version)
}
