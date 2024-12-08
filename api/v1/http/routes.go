package http

import (
	infra "github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
	"net/http"
)

var (
	GroupV1  = infra.NewGroup("api").Group("v1")
	GroupV1P = GroupV1.Group("p")

	// Gateway
	GatewayAlive       = GroupV1.NewEndpoint("/alive", http.MethodGet)
	GatewayPublicAlive = GroupV1P.NewEndpoint("/alive/gateway", http.MethodGet)

	// Docs
	GroupV1Documentation = GroupV1P.Group("documentation")
	SwaggerDocs          = GroupV1Documentation.NewEndpoint("/:service/docs", http.MethodGet)
	SwaggerSwagger       = GroupV1Documentation.NewEndpoint("/:service/swagger", http.MethodGet)
	SwaggerJson          = GroupV1Documentation.NewEndpoint("/:service/swagger.json", http.MethodGet)

	// Auth
	GroupV1Auth                         = GroupV1.Group("auth")
	GroupV1PAuth                        = GroupV1P.Group("auth")
	GetTokenInternalProviders           = GroupV1PAuth.NewEndpoint("/:provider/login", http.MethodPost)
	LoginByExternalProvider             = GroupV1PAuth.NewEndpoint("/:provider/login", http.MethodGet)
	GetTokenByCallBackExternalProviders = GroupV1PAuth.NewEndpoint("/:provider/callback", http.MethodGet)
	SignupInternalProviders             = GroupV1PAuth.NewEndpoint("/:provider/signup", http.MethodPost)
	SignupInternalProvidersConfirmation = GroupV1PAuth.NewEndpoint("/:provider/signup/confirmation", http.MethodPost)
	CreatePassword                      = GroupV1Auth.NewEndpoint("/signup/create-password", http.MethodPost)
	Logout                              = GroupV1PAuth.NewEndpoint("/logout", http.MethodPost)
	Refresh                             = GroupV1PAuth.NewEndpoint("/refresh-token", http.MethodPost)

	// User
	GroupV1User  = GroupV1.Group("user")
	GroupV1PUser = GroupV1P.Group("user")
	GetUserMe    = GroupV1User.NewEndpoint("/me", http.MethodGet)
	Onboarding   = GroupV1PUser.NewEndpoint("/onboarding", http.MethodPost)

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
