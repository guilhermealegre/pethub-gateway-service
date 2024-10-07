package v1

import (
	v1 "bitbucket.org/asadventure/be-gateway-service/internal/access/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
)

const (
	Message = "success"
)

type Model struct {
	app domain.IApp
}

func NewModel(app domain.IApp) v1.IModel {
	return &Model{
		app: app,
	}
}

func (m *Model) Get(ctx domain.IContext) (_ *v1.AccessClearance, err error) {
	obj := &v1.AccessClearance{
		Message: Message,
	}

	return obj, nil
}
