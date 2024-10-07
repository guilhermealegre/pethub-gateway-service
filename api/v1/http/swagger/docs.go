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
	_ "bitbucket.org/asadventure/be-gateway-service/internal/access/controller/v1"   // alive controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/alive/controller/v1"    // access controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/customer/controller/v1" // customer controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/logging/controller/v1"  // logging controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/order/controller/v1"    // order controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/store/controller/v1"    // store controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/swagger/controller/v1"  // swagger controller
	_ "bitbucket.org/asadventure/be-gateway-service/internal/user/controller/v1"     // auth controller
)
