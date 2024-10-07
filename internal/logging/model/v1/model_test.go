package v1

import (
	"testing"

	v1Streaming "bitbucket.org/asadventure/be-gateway-service/internal/logging/streaming/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/app"
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
