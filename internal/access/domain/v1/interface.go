package v1

import (
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
)

type IModel interface {
	Get(ctx domain.IContext) (*AccessClearance, error)
}
