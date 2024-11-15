/*
	 Gateway Service

	 # Gateway Service API

	 Schemes: http, https
	 BasePath: /v1
	 Version: 1.0

	 Consumes:
	 - application/json

	 Produces:
	 - application/json

	 SecurityDefinitions:
		Bearer:
		  type: apiKey
		  name: Authorization
		  in: header

	 swagger:meta
*/
package swagger

import (
	_ "github.com/guilhermealegre/pethub-gateway-service/internal/alive/controller/v1"   // access controller
	_ "github.com/guilhermealegre/pethub-gateway-service/internal/logging/controller/v1" // logging controller
	_ "github.com/guilhermealegre/pethub-gateway-service/internal/swagger/controller/v1" // swagger controller
	_ "github.com/guilhermealegre/pethub-gateway-service/internal/user/controller/v1"    // auth controller
)
