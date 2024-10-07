package response

import "bitbucket.org/asadventure/be-core-lib/response"

// swagger:model SwaggerAccessClearanceResponse
type swaggerAccessClearanceResponse struct { //nolint:all
	response.Response
	Data AccessClearanceResponse `json:"data"`
}

// swagger:model AccessClearanceResponse
type AccessClearanceResponse struct {
	Message string `json:"message"`
}
