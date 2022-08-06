// The MiddlewareComponent interface describes the functionality required by a middleware component.
// The Init method is used to perform any one-off setup,
// and the other method, named ProcessRequest, is responsible for processing HTTP requests.
// The parameters defined by the ProcessRequest method are a pointer to a ComponentContext struct
// and a function that passes the request to the next component in the pipeline.
// Everything a component needs to process a request is provided by the ComponentContext struct,
// through which http.Request and http.ResponseWriter can be accessed.
// The ComponentContext struct also defines an unexported error field, which is used to indicate a problem processing a request
// and which is set using the Error method.
package pipeline

import (
	"net/http"
)

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (mwc *ComponentContext) Error(err error) {
	mwc.error = err
}

func (mwc *ComponentContext) GetError() error {
	return mwc.error
}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentContext, next func(*ComponentContext))
}

type ServicesMiddlwareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
