package v1

import "bitbucket.org/asadventure/be-gateway-service/api/v1/http/envelope/response"

func (pa *AccessClearance) FromDomainToApi() *response.AccessClearanceResponse {
	if pa == nil {
		return nil
	}

	return &response.AccessClearanceResponse{
		Message: pa.Message,
	}
}
