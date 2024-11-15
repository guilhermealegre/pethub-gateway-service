package main

import (
	netHttp "net/http"

	logging "bitbucket.org/asadventure/be-logging-service/api/v1/grpc/logging_service_logging"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/grpc"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/logger/writer"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/meter"
	grpcInfra "github.com/guilhermealegre/pethub-gateway-service/internal/infrastructure/grpc"
	loggingModel "github.com/guilhermealegre/pethub-gateway-service/internal/logging/model/v1"
	loggingStreaming "github.com/guilhermealegre/pethub-gateway-service/internal/logging/streaming/v1"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/tracer"

	v1AliveController "github.com/guilhermealegre/pethub-gateway-service/internal/alive/controller/v1"
	v1AliveModel "github.com/guilhermealegre/pethub-gateway-service/internal/alive/model/v1"
	v1LoggingController "github.com/guilhermealegre/pethub-gateway-service/internal/logging/controller/v1"
	v1Middleware "github.com/guilhermealegre/pethub-gateway-service/internal/middleware/v1"
	_ "github.com/guilhermealegre/pethub-gateway-service/internal/request/config"
	v1RequestModel "github.com/guilhermealegre/pethub-gateway-service/internal/request/model/v1"
	v1SwaggerController "github.com/guilhermealegre/pethub-gateway-service/internal/swagger/controller/v1"
	v1UploaderController "github.com/guilhermealegre/pethub-gateway-service/internal/uploader/controller/v1"
	v1UserController "github.com/guilhermealegre/pethub-gateway-service/internal/user/controller/v1"

	"os"

	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/app"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/http"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/logger"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/sqs"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/validator"
	_ "github.com/lib/pq" // postgres driver
)

func main() {
	// app initialization
	newApp := app.New(nil)
	newHttp := http.New(newApp, nil)
	newValidation := validator.New(newApp).
		AddFieldValidators().
		AddStructValidators()
	newGrpc := grpc.New(newApp, nil)

	loggingGrpcClient := newGrpc.GetClient(grpcInfra.LoggingClient)
	loggingClient := logging.NewLoggingClient(loggingGrpcClient)
	loggingModel := loggingModel.NewModel(newApp, loggingStreaming.NewStreaming(newApp, loggingClient))
	newLogger := logger.New(newApp, nil, writer.NewGeneric(loggingModel.Log, nil))

	newTracer := tracer.New(newApp, nil)
	newMeter := meter.New(newApp, nil)
	newSQS := sqs.New(newApp, nil)

	// models
	aliveModel := v1AliveModel.NewModel(newApp)
	accessModel := v1AccessModel.NewModel(newApp)
	requestModel := v1RequestModel.NewModel(newApp, netHttp.DefaultClient)

	newHttp.
		//middlewares
		WithMiddleware(v1Middleware.NewPrintRequestMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewAuthorizationMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewIncreaseTTLMiddleware(newApp)).
		WithMiddleware(v1Middleware.NewApiKeyMiddleware(newApp)).
		//controllers
		WithController(v1SwaggerController.NewController(newApp)).
		WithController(v1AliveController.NewController(newApp, aliveModel)).
		WithController(v1UserController.NewController(newApp, requestModel)).
		WithController(v1UploaderController.NewController(newApp, requestModel)).
		WithController(v1LoggingController.NewController(newApp, requestModel))

	newApp.
		WithLogger(newLogger).
		WithSQS(newSQS).
		WithTracer(newTracer).
		WithMeter(newMeter).
		WithValidator(newValidation).
		WithHttp(newHttp).
		WithGrpc(newGrpc)

	// start app
	if err := newApp.Start(); err != nil {
		os.Exit(1)
	}
}
