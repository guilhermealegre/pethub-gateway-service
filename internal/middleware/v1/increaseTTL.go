package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type IncreaseTTLMiddleware struct {
	app domain.IApp
}

func NewIncreaseTTLMiddleware(app domain.IApp) domain.IMiddleware {
	return &IncreaseTTLMiddleware{
		app: app,
	}
}

func (c *IncreaseTTLMiddleware) RegisterMiddlewares() {

}

func (c *IncreaseTTLMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}
