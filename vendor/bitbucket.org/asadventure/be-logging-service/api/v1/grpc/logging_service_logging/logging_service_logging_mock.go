package logging_service_logging

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LoggingClientMock struct {
	mock.Mock
}

func NewLoggingClientMock() *LoggingClientMock {
	return &LoggingClientMock{}
}

func (c *LoggingClientMock) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	var optsList = []interface{}{ctx, in}
	for _, v := range opts {
		optsList = append(optsList, v)
	}

	args := c.Called(optsList...)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}
