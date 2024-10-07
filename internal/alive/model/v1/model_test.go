package v1

import (
	"os"
	"strconv"
	"testing"

	v1 "bitbucket.org/asadventure/be-gateway-service/internal/alive/domain/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/app"
	appConfig "bitbucket.org/asadventure/be-infrastructure-lib/app/config"
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	httpLib "bitbucket.org/asadventure/be-infrastructure-lib/http"
	httpConfig "bitbucket.org/asadventure/be-infrastructure-lib/http/config"
	"github.com/stretchr/testify/assert"
)

// TestStoreModelAlive test for the alive method
func TestStoreModelAlive(t *testing.T) {
	hostName, _ := os.Hostname()

	aliveResponse := &v1.Alive{
		ServerName: "gateway",
		Port:       "80",
		Hostname:   hostName,
		Message:    "I AM ALIVE!!!",
	}

	testCases := []*TestCase{
		testCaseAlive(),
	}

	newHttp := httpLib.NewHttpMock()
	port, _ := strconv.Atoi(aliveResponse.Port)
	newHttp.On("Config").Return(&httpConfig.Config{
		Host: aliveResponse.Hostname,
		Port: port,
	})

	newApp := app.NewAppMock()
	newApp.On("Config").Return(&appConfig.Config{
		Name: aliveResponse.ServerName,
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

// TestStoreModelPublicAlive test for the alive method
func TestStoreModelPublicAlive(t *testing.T) {
	aliveResponse := &v1.PublicAlive{
		Name:    "gateway",
		Message: "I AM ALIVE!!!",
	}

	testCases := []*TestCase{
		testCasePublicAlive(),
	}

	newHttp := httpLib.NewHttpMock()
	hostName, _ := os.Hostname()
	newHttp.On("Config").Return(&httpConfig.Config{
		Host: hostName,
		Port: 80,
	})

	newApp := app.NewAppMock()
	newApp.On("Config").Return(&appConfig.Config{
		Name: aliveResponse.Name,
		Env:  "local",
	})
	newApp.On("Http").Return(newHttp)

	for _, test := range testCases {
		test.Log(t)

		// model
		model := NewModel(newApp)
		result, err := model.GetPublic(test.Arguments[0].(domain.IContext))

		assert.Equal(t, test.Expected[1] == nil, err == nil)    // check nil error
		assert.Equal(t, test.Expected[0] == nil, result == nil) // check nil result
		if test.Expected[0] != nil {
			assert.Equal(t, test.Expected[0], result) // check result object
		}
	}
}