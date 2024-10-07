package v1

import (
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
)

type IController interface {
	domain.IController
	Redirect(gCtx *gin.Context)
}

type IModel interface {
	Log(message []byte) error
}

type IStreaming interface {
	Log(message []byte) error
}
