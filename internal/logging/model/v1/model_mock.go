package v1

import (
	"github.com/stretchr/testify/mock"
)

func NewModelMock() *ModelMock {
	return &ModelMock{}
}

type ModelMock struct {
	mock.Mock
}

func (m *ModelMock) Log(message []byte) error {
	args := m.Called(message)
	return args.Error(0)
}
