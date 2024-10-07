package main

import (
	netHttp "net/http"

	grpcInfra "bitbucket.org/asadventure/be-gateway-service/internal/infrastructure/grpc"
	loggingModel "bitbucket.org/asadventure/be-gateway-service/internal/logging/model/v1"
	loggingStreaming "bitbucket.org/asadventure/be-gateway-service/internal/logging/streaming/v1"
	"bitbucket.org/asadventure/be-infrastructure-lib/grpc"
	"bitbucket.org/asadventure/be-infrastructure-lib/logger/writer"
	"bitbucket.org/asadventure/be-infrastructure-lib/meter"
	logging "bitbucket.org/asadventure/be-logging-service/api/v1/grpc/logging_service_logging"

	"bitbucket.org/asadventure/be-infrastructure-lib/tracer"

	v1AccessController "bitbucket.org/asadventure/be-gateway-service/internal/access/controller/v1"
	v1AccessModel "bitbucket.org/asadventure/be-gateway-service/internal/access/model/v1"
	v1AliveController "bitbucket.org/asadventure/be-gateway-service/internal/alive/controller/v1"
	v1AliveModel "bitbucket.org/asadventure/be-gateway-service/internal/alive/model/v1"
	v1CustomerController "bitbucket.org/asadventure/be-gateway-service/internal/customer/controller/v1"
	v1FallbackController "bitbucket.org/asadventure/be-gateway-service/internal/fallback/controller/v1"
	v1LoggingController "bitbucket.org/asadventure/be-gateway-service/internal/logging/controller/v1"
	v1Middleware "bitbucket.org/asadventure/be-gateway-service/internal/middleware/v1"
	v1OrderController "bitbucket.org/asadventure/be-gateway-service/internal/order/controller/v1"
	_ "bitbucket.org/asadventure/be-gateway-service/internal/request/config"
	v1RequestModel "bitbucket.org/asadventure/be-gateway-service/internal/request/model/v1"
	v1StoreController "bitbucket.org/asadventure/be-gateway-service/internal/store/controller/v1"
	v1SwaggerController "bitbucket.org/asadventure/be-gateway-service/internal/swagger/controller/v1"
	v1UploaderController "bitbucket.org/asadventure/be-gateway-service/internal/uploader/controller/v1"
	v1UserController "bitbucket.org/asadventure/be-gateway-service/internal/user/controller/v1"

	"os"

	"bitbucket.org/asadventure/be-infrastructure-lib/app"
	"bitbucket.org/asadventure/be-infrastructure-lib/http"
	"bitbucket.org/asadventure/be-infrastructure-lib/logger"
	"bitbucket.org/asadventure/be-infrastructure-lib/sqs"
	"bitbucket.org/asadventure/be-infrastructure-lib/validator"
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
		WithController(v1AccessController.NewController(newApp, accessModel)).
		WithController(v1UserController.NewController(newApp, requestModel)).
		WithController(v1StoreController.NewController(newApp, requestModel)).
		WithController(v1CustomerController.NewController(newApp, requestModel)).
		WithController(v1OrderController.NewController(newApp, requestModel)).
		WithController(v1UploaderController.NewController(newApp, requestModel)).
		WithController(v1LoggingController.NewController(newApp, requestModel)).
		WithController(v1FallbackController.NewController(newApp, requestModel))

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
