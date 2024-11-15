package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"

	http2 "github.com/guilhermealegre/pethub-gateway-service/api/v1/http"
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
	http2.GroupV1User.AddMiddleware(c)
	http2.GroupV1Customer.AddMiddleware(c)
	http2.GroupV1Store.AddMiddleware(c)
	http2.GroupV1Order.AddMiddleware(c)
	http2.GroupV1Uploader.AddMiddleware(c)
}

func (c *IncreaseTTLMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		auth.IncreaseActivityTTLInXMinutes,
	}
}
