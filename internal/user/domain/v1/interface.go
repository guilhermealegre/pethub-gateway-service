package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/domain"
)

type IController interface {
	domain.IController
	Redirect(gCtx *gin.Context)
}
