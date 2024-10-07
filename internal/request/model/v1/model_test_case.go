package v1

import (
	"net/http"

	"bitbucket.org/asadventure/be-core-lib/test"
)

type TestCase struct {
	test.BaseTestCase
	Mock test.MapCall
}

func testCaseRequestModel() *TestCase {
	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Getting customer by id",
			Call: test.Call{
				Arguments: []interface{}{http.MethodGet, "http://localhost:8083/v1/user", nil},
				Expected:  []interface{}{},
				Other: test.MapSomething{
					"query": map[string]string{
						"id":   "5",
						"name": "joe",
					},
				},
			},
		},
		Mock: test.MapCall{
			"status_code": test.CallList{
				test.Call{
					Arguments: []interface{}{http.StatusOK},
					Expected:  []interface{}{http.StatusOK},
				},
			},
			"header": test.CallList{
				test.Call{
					Arguments: []interface{}{
						map[string]string{
							"Content-Type": "application/json",
						},
					},
					Expected: []interface{}{
						map[string]string{
							"Content-Type": "application/json",
						},
					},
				},
			},
		},
	}
}
