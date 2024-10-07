package sqs

import (
	"bitbucket.org/asadventure/be-infrastructure-lib/domain"
	sqsConfig "bitbucket.org/asadventure/be-infrastructure-lib/sqs/config"
	middlewares "bitbucket.org/asadventure/be-infrastructure-lib/sqs/middlewares"
	"github.com/aws/aws-sdk-go/service/sqs"
	sqsSdk "github.com/aws/aws-sdk-go/service/sqs"
)

// SQS service
type SQS struct {
	// Name
	name string
	// App
	app domain.IApp
	// Configuration
	config *sqsConfig.Config
	// Connections
	connections map[string]*Connection
	// Additional Config Type
	additionalConfigType interface{}
	// Started
	started bool
}

type Connection struct {
	// Service Name
	serviceName string
	// Name
	name string
	// App
	app domain.IApp
	// Config
	config *sqsConfig.Connection
	// Consumer
	consumer *sqs.SQS
	// Producer Connection
	producer *sqs.SQS
	// Migration
	migration *Migration
	// Consumers
	consumers []domain.ISQSConsumer
	// Middlewares
	middlewares []middlewares.Middleware
}

type Migration struct {
	name       string
	config     *sqsConfig.Connection
	connection *sqsSdk.SQS
}
