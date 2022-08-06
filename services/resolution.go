// set of features includes the functions that allow services to be resolved
package services

import (
	"context"
	"errors"
	"reflect"
)

// For convenience, the GetService function resolves a service using the background context.
func GetService(target interface{}) error {
	return GetServiceForContext(context.Background(), target)
}

// The GetServiceForContext accepts a context and a pointer to a value that can be set using reflection.
func GetServiceForContext(c context.Context, target interface{}) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = resolveServiceFromValue(c, targetValue)
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
