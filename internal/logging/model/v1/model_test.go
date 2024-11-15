package v1

import (
	"testing"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	v1Streaming "github.com/guilhermealegre/pethub-gateway-service/internal/logging/streaming/v1"
	"github.com/stretchr/testify/assert"
)

func TestLoggingModelLog(t *testing.T) {
	testCases := []*TestCase{
		testCaseLoggingModelLogWithSuccess(),
		testCaseLoggingModelLogWithError(),
	}

	newApp := app.NewAppMock()

	for _, test := range testCases {
		test.Log(t)

		// streaming
		streaming := v1Streaming.NewStreamingMock()
		test.Streaming.Setup(streaming)

		// model
		model := NewModel(newApp, streaming)
		err := model.Log(test.Arguments[0].([]byte))

		assert.Equal(t, test.Expected[0] == nil, err == nil) // check nil error
	}
}
