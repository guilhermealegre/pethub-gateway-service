package v1

import (
	http2 "bitbucket.org/asadventure/be-gateway-service/api/v1/http"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain/auth"
	"github.com/gin-gonic/gin"
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
	http2.GroupV1Customer.AddMiddleware(c)
	http2.GroupV1Store.AddMiddleware(c)
	http2.GroupV1Order.AddMiddleware(c)
	http2.GroupV1Uploader.AddMiddleware(c)
}

func (c *AuthorizationMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		auth.BuildAuthorizationHeader,
	}
}
