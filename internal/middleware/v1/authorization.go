package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
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

}

func (c *AuthorizationMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//auth.BuildAuthorizationHeader,
	}
}
