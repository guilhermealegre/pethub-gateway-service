package v1

import (
	"os"

	"bitbucket.org/asadventure/be-core-lib/test"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/context"
	v1 "github.com/guilhermealegre/pethub-gateway-service/internal/alive/domain/v1"
)

type TestCase struct {
	test.BaseTestCase
}

func testCaseAlive() *TestCase {
	hostName, _ := os.Hostname()

	aliveResponse := &v1.Alive{
		ServerName: "gateway",
		Port:       "80",
		Hostname:   hostName,
		Message:    "I AM ALIVE!!!",
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Getting alive",
			Call: test.Call{
				Arguments: []interface{}{&context.Context{}},
				Expected:  []interface{}{aliveResponse, nil},
			},
		},
	}
}

func testCasePublicAlive() *TestCase {
	aliveResponse := &v1.PublicAlive{
		Name:    "gateway",
		Message: "I AM ALIVE!!!",
	}

	return &TestCase{
		BaseTestCase: test.BaseTestCase{
			Test:        1,
			Description: "Getting alive",
			Call: test.Call{
				Arguments: []interface{}{&context.Context{}},
				Expected:  []interface{}{aliveResponse, nil},
			},
		},
	}
}
