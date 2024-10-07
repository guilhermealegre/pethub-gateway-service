package sqs

import (
	"context"

	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/mock"
)

func NewConnectionMock() *ConnectionMock {
	return &ConnectionMock{}
}

type ConnectionMock struct {
	mock.Mock
}

func (s *ConnectionMock) Connect() error {
	args := s.Called()
	return args.Error(0)
}

func (s *ConnectionMock) Produce(ctx context.Context, queue string, messageAttributes map[string]*sqs.MessageAttributeValue, messages ...string) error {
	params := []interface{}{ctx, queue, messageAttributes}
	for _, p := range messages {
		params = append(params, p)
	}
	args := s.Called(params...)
	return args.Error(0)
}

func (s *ConnectionMock) Consume(maskedQueue string, consumer domain.ISQSConsumer) {
	s.Called(maskedQueue, consumer)
}
