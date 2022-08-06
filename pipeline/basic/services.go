/*
 This middleware component modifies the Context associated with the request so that context-scoped services can be used during request processing.
 The http.Request.Context method is used to get the standard Context created with the request,
 which is prepared for services and then updated using the WithContext method.

Once the context has been prepared, the request is passed along the pipeline by invoking the function received through the parameter named next:
     ... next(ctx) ...
This parameter gives middleware components control over request processing and allows it to modify the context data that subsequent components receive.
It also allows components to short-circuit request processing by not invoking the next function.
*/
package basic

import (
	"platform/pipeline"
	"platform/services"
)

type ServicesComponent struct{}

func (c *ServicesComponent) Init() {}

func (c *ServicesComponent) ProcessRequest(ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext)) {
	reqContext := ctx.Request.Context()
	ctx.Request.WithContext(services.NewServiceContext(reqContext))
	next(ctx)
}
