package v1

import (
	"github.com/stretchr/testify/mock"
)

func NewStreamingMock() *StreamingMock {
	return &StreamingMock{}
}

type StreamingMock struct {
	mock.Mock
}

func (m *StreamingMock) Log(message []byte) error {
	args := m.Called(message)
	return args.Error(0)
}
