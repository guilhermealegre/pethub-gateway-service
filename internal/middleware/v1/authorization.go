package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"

	http2 "github.com/guilhermealegre/pethub-gateway-service/api/v1/http"
)

type AuthorizationMiddleware struct {
	app domain.IApp
}

func NewAuthorizationMiddleware(app domain.IApp) domain.IMiddleware {
	return &AuthorizationMiddleware{
		app: app,
	}
}

func (c *AuthorizationMiddleware) RegisterMiddlewares() {
	http2.GroupV1User.AddMiddleware(c)
	http2.GroupV1Uploader.AddMiddleware(c)
}

func (c *AuthorizationMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//auth.BuildAuthorizationHeader,
	}
}
