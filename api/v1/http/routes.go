package http

import (
	"net/http"
)

var (
	GroupV1  = infra.NewGroup("v1")
	GroupV1P = GroupV1.Group("p")

	// Gateway
	GatewayAlive       = GroupV1.NewEndpoint("/alive", http.MethodGet)
	GatewayPublicAlive = GroupV1P.NewEndpoint("/alive/gateway", http.MethodGet)

	// Docs
	GroupV1Documentation = GroupV1P.Group("documentation")
	SwaggerDocs          = GroupV1Documentation.NewEndpoint("/:service/docs", http.MethodGet)
	SwaggerSwagger       = GroupV1Documentation.NewEndpoint("/:service/swagger", http.MethodGet)
	SwaggerJson          = GroupV1Documentation.NewEndpoint("/:service/swagger.json", http.MethodGet)

	// Fallback
	PublicAlive = GroupV1P.NewEndpoint("/alive/:service", http.MethodGet)

	// Uploader
	GroupV1Uploader  = GroupV1.Group("uploader")
	GroupV1PUploader = GroupV1P.Group("uploader")
	ImageUpload      = GroupV1Uploader.NewEndpoint("/image/upload", http.MethodPost)
	ImagePUpload     = GroupV1PUploader.NewEndpoint("/image/upload", http.MethodPost)

	// Logging
	GroupV1Logging     = GroupV1.Group("logging")
	GroupV1PLogging    = GroupV1P.Group("logging")
	LoggingCreateFeLog = GroupV1PLogging.NewEndpoint("/log", http.MethodPost)
)
