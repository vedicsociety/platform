// Create enhancements that make service resolution simpler and easier.
// add support for executing functions directly
/*
   example:
   func main {
        services.Call(writeMessage)
        }
   The function is passed to Call, which will inspect its parameters and resolve them using services.
   (Note that parentheses do not follow the function name because that would invoke the function rather than pass it to services.Call.)
   We no longer have to request services directly and can rely on the services package to take care of the details.
*/
package services

import (
	"context"
	"errors"
	"reflect"
)

// The Call function is a convenience for use when a Context isn’t available.
// The implementation of this feature relies on the code used to invoke factory functions.
func Call(target interface{}, otherArgs ...interface{}) ([]interface{}, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

// The CallForContext function receives a function and uses services to produce the values that are used as arguments to invoke the function.
func CallForContext(c context.Context, target interface{}, otherArgs ...interface{}) (results []interface{}, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultVals := invokeFunction(c, targetValue, otherArgs...)
		results = make([]interface{}, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			results[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("Only functions can be invoked")
	}
	return
}
