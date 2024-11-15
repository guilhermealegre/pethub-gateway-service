package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
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
