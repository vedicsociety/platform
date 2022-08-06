// We implement the Scoped lifecycle using the context package in the standard library, which I described in Chapter 30.
// A Context will be automatically created for each HTTP request received by the server,
// which means that all the request handling code that processes that request can share the same set of services so that,
// for example, a single struct that provides session information can be used throughout processing for a given request.
// This make it easier to work with contexts
package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

// The NewServiceContext function derives a context using the WithValue function, adding a map that stores the services that have been resolved.
// See Chapter 30 for details of the different ways that contexts can be derived.
func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(serviceMap))
	} else {
		return c
	}
}
