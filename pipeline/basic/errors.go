/*
The request pipeline allows components to terminate processing when an error arises.
This component recovers from any panic that occurs when subsequent components process the request and also handles any expected error.
In both cases, the error is logged, and the response status code is set to indicate an error.
*/
package basic

import (
	"fmt"
	"net/http"

	"github.com/vedicsociety/platform/logging"
	"github.com/vedicsociety/platform/pipeline"
	"github.com/vedicsociety/platform/services"
)

type ErrorComponent struct{}

func recoveryFunc(ctx *pipeline.ComponentContext, logger logging.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %v", fmt.Sprint(arg))
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *ErrorComponent) Init() {}

func (c *ErrorComponent) ProcessRequest(ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext)) {

	var logger logging.Logger
	services.GetServiceForContext(ctx.Context(), &logger)
	defer recoveryFunc(ctx, logger)
	next(ctx)
	if ctx.GetError() != nil {
		logger.Debugf("Error: %v", ctx.GetError())
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
