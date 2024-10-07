package v1

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	httpCore "bitbucket.org/asadventure/be-core-lib/http"
	"bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	"bitbucket.org/asadventure/be-infrastructure-lib/app"
	"bitbucket.org/asadventure/be-infrastructure-lib/context"
	httpLib "bitbucket.org/asadventure/be-infrastructure-lib/http"
	"bitbucket.org/asadventure/be-infrastructure-lib/logger"
	"bitbucket.org/asadventure/be-infrastructure-lib/logger/logging"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestRequestModel test for the request model
func TestRequestModel(t *testing.T) {
	// app
	newApp := app.New(nil)

	// mock logger (on error scenario)
	logMock := logging.NewLoggingMock()
	logMock.On("Do", mock.Anything)

	loggerMock := logger.NewLoggerMock()
	loggerMock.On("Log").Return(logMock)

	testCases := []*TestCase{
		testCaseRequestModel(),
	}

	for _, test := range testCases {
		test.Log(t)

		var body string
		if len(test.Call.Expected) > 0 {
			body = test.Call.Expected[0].(string)
		}
		httpResponse := &http.Response{
			StatusCode: test.Mock["status_code"][0].Arguments[0].(int),
			Body:       io.NopCloser(strings.NewReader(body)),
		}

		// http test
		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		// ctx
		ctx := context.NewContextMock()

		// http
		newApp.
			WithLogger(loggerMock).
			WithHttp(
				httpLib.New(newApp, nil).
					// router
					WithRouter(router),
			)

		// new request object
		body = ""
		if len(test.Arguments) > 2 && test.Arguments[2] != nil {
			body = test.Arguments[2].(string)
		}
		req, err := http.NewRequest(test.Arguments[0].(string), test.Arguments[1].(string), strings.NewReader(body))
		assert.NoError(t, err)

		ctx.On("Request").Return(req)

		for key, value := range test.Mock["header"][0].Arguments[0].(map[string]string) {
			req.Header.Add(key, value)
		}

		expectedQueryParams := url.Values{}
		for key, value := range test.Call.Other["query"].(map[string]string) {
			expectedQueryParams.Set(key, value)
		}

		req.URL.RawQuery = expectedQueryParams.Encode()
		httpResponse.Request = req
		httpResponse.Header = req.Header

		// model
		httpClientMock := httpCore.NewHttpClientMock()
		httpClientMock.On("Do", req).Return(httpResponse, nil)
		model := NewModel(newApp, httpClientMock)

		// Test the implementation
		resp, _ := model.Redirect(ctx, &config.Endpoint{
			Protocol: "http",
			Host:     "localhost",
			Port:     "8080",
		})

		// assert the response status code is correct
		assert.Equal(t, test.Mock["status_code"][0].Expected[0].(int), resp.StatusCode)

		// assert the response headers are correct
		for key, value := range test.Mock["header"][0].Expected[0].(map[string]string) {
			assert.Equal(t, value, resp.Header.Get(key))
		}

		// assert the response query parameters are correct
		for key, values := range expectedQueryParams {
			for _, value := range values {
				assert.Equal(t, value, resp.Request.URL.Query().Get(key))
			}
		}

		if len(test.Expected) > 0 {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			assert.Nil(t, err)
			assert.Equal(t, test.Expected[0].(string), string(body)) // check result object
		}
	}
}
