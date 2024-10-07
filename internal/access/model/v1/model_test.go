package v1

import (
	"testing"

	"bitbucket.org/asadventure/be-core-lib/test"

	httpLib "bitbucket.org/asadventure/be-infrastructure-lib/http"
	httpConfig "bitbucket.org/asadventure/be-infrastructure-lib/http/config"

	v1 "bitbucket.org/asadventure/be-gateway-service/internal/access/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/app"
	appConfig "bitbucket.org/asadventure/be-infrastructure-lib/app/config"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	test.BaseTestCase
	test.Call
}

// TestStoreModelAlive test for the alive method
func TestStoreModelAccessClearance(t *testing.T) {

	aliveResponse := &v1.AccessClearance{
		Message: "success",
	}

	testCases := []*TestCase{
		{
			BaseTestCase: test.BaseTestCase{
				Test:        1,
				Description: "Getting Access Clearance",
			},
			Call: test.Call{
				Arguments: []interface{}{&context.Context{}},
				Expected:  []interface{}{aliveResponse, nil},
			},
		},
	}
	newHttp := httpLib.NewHttpMock()
	newHttp.On("Config").Return(&httpConfig.Config{
		Host: "localhost",
		Port: 8080,
	})

	newApp := app.NewAppMock()
	newApp.On("Config").Return(&appConfig.Config{
		Name: "gateway",
		Env:  "local",
	})
	newApp.On("Http").Return(newHttp)

	for _, test := range testCases {
		test.Log(t)

		// model
		model := NewModel(newApp)
		result, err := model.Get(test.Arguments[0].(domain.IContext))

		assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
		assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result) // check result object
		}
	}
}
