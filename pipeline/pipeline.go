/* 	Create the pipeline that will process requests.
The CreatePipeline function is the most important part of this listing because it accepts a series of components
and connects them to produce a function that accepts a pointer to a ComponentContext struct.
This function invokes the ProcessRequest method of the first component in the pipeline with a next argument
 that invokes the ProcessRequest method of the next component.
 This chain passes on the ComponentContext struct to all the components in turn, unless one of them calls the Error method.
 Requests are processed using the ProcessRequest method, which creates the ComponentContext value and uses it to start request handling.

 The definition of the component interface and pipeline are simple, but they provide a flexible foundation on which components can be written.
 Applications can define and choose their own components, but there are some basic features that defined in the basic folder.
*/

package pipeline

import (
	"net/http"
	"github.com/tsiparinda/platform/services"
	"reflect"
)

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) { /* do nothing */ }

func CreatePipeline(components ...interface{}) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		currentComponent := components[i]
		services.Populate(currentComponent)
		nextFunc := f
		if servComp, ok := currentComponent.(ServicesMiddlwareComponent); ok {
			f = createServiceDependentFunction(currentComponent, nextFunc)
			servComp.Init()
		} else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
			f = func(context *ComponentContext) {
				if context.error == nil {
					stdComp.ProcessRequest(context, nextFunc)
				}
			}
			stdComp.Init()
		} else {
			panic("Value is not a middleware component")
		}
	}
	return f
}

func createServiceDependentFunction(component interface{},
	nextFunc RequestPipeline) RequestPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if method.IsValid() {
		return func(context *ComponentContext) {
			if context.error == nil {
				_, err := services.CallForContext(context.Request.Context(),
					method.Interface(), context, nextFunc)
				if err != nil {
					context.Error(err)
				}
			}
		}
	} else {
		panic("No ProcessRequestWithServices method defined")
	}
}

func (pl RequestPipeline) ProcessRequest(req *http.Request,
	resp http.ResponseWriter) error {
	deferredWriter := &DeferredResponseWriter{ResponseWriter: resp}
	ctx := ComponentContext{
		Request:        req,
		ResponseWriter: deferredWriter,
	}
	pl(&ctx)
	if ctx.error == nil {
		deferredWriter.FlushData()
	}
	return ctx.error
}
