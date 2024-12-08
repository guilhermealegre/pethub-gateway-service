package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

const (
	ApiKey = "apikey"
)

type ApiKeyMiddleware struct {
	app domain.IApp
}

func NewApiKeyMiddleware(app domain.IApp) domain.IMiddleware {
	return &ApiKeyMiddleware{
		app: app,
	}
}

func (c *ApiKeyMiddleware) RegisterMiddlewares() {

}

func (c *ApiKeyMiddleware) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		c.checkApiKey,
	}
}

func (c *ApiKeyMiddleware) checkApiKey(ctx *gin.Context) {
	/*
		// validate authorization header
		authorizationString := ctx.GetHeader(auth.HeaderAuthorization)
		if authorizationString == "" {
			// return unauthorized error
			ctx.JSON(http.StatusUnauthorized, errors.ErrorAuthorizationMissing())
			ctx.Abort()
			return
		}

		var apiKeyStr string
		if aToken := strings.Split(authorizationString, " "); len(aToken) == 2 {
			if helpers.TrimAndLowerStr(aToken[0]) != ApiKey {
				// return unauthorized error
				ctx.JSON(http.StatusUnauthorized, errors.ErrorAuthorizationMissing())
				ctx.Abort()
				return
			}
			apiKeyStr = aToken[1]
		} else {
			// return unauthorized error
			ctx.JSON(http.StatusUnauthorized, errors.ErrorAuthorizationMissing())
			ctx.Abort()
			return
		}

		for _, key := range c.app.Http().Config().ApiKeys {
			if key == apiKeyStr {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, errors.ErrorInvalidApiKey())
		ctx.Abort()

	*/
}
