package v1

import (
	"bitbucket.org/asadventure/be-core-lib/test"
	errorCodes "bitbucket.org/asadventure/be-infrastructure-lib/errors"
)

type TestCase struct {
	test.BaseTestCase
	Streaming test.MapCall
}

func testCaseLoggingModelLogWithSuccess() *TestCase {
	message := []byte("hello world")

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Log with success",
			Call: test.Call{
				Arguments: []interface{}{message},
				Expected:  []interface{}{nil},
			},
		},
		Streaming: test.MapCall{
			"Log": test.CallList{
				test.Call{
					Arguments: []interface{}{message},
					Expected:  []interface{}{nil},
				},
			},
		},
	}
}

func testCaseLoggingModelLogWithError() *TestCase {
	message := []byte("hello world")
	err := errorCodes.ErrorFieldEmpty().Formats("message")

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        2,
			Description: "Log with error",
			Call: test.Call{
				Arguments: []interface{}{message},
				Expected:  []interface{}{err},
			},
		},
		Streaming: test.MapCall{
			"Log": test.CallList{
				test.Call{
					Arguments: []interface{}{message},
					Expected:  []interface{}{err},
				},
			},
		},
	}
}
