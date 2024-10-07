package v1

import (
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/gin-gonic/gin"
)

type IController interface {
	GetPublic(ctx *gin.Context)
}

type IModel interface {
	Get(ctx domain.IContext) (*Alive, error)
	GetPublic(ctx domain.IContext) (*PublicAlive, error)
}
