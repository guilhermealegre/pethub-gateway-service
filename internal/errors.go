package internal

import (
	"bitbucket.org/asadventure/be-core-lib/errors"
	"bitbucket.org/asadventure/be-infrastructure-lib/errors/config"
)

// Generic error codes
var (
	ErrorGeneric            = config.GetError("1", "%s", errors.Error)
	ErrorInvalidInputFields = config.GetError("2", "The field: %s as invalid value: %v", errors.Info)
)
