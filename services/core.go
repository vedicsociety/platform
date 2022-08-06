// We use handle service registration by inspecting a factory function and using its result to determine the interface that it handles.
// None of the functions are exported. see the registration.go in the services folder
package services

import (
	"context"
	"fmt"
	"reflect"
)

// The BindingMap struct represents the combination of a factory function, expressed as a reflect.Value, and a lifecycle.
type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

// The addService function is used to register a service, which it does by creating BindingMap and adding to the map
// assigned to the services variable.
func addService(life lifecycle, factoryFunc interface{}) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		services[factoryFuncType.Out(0)] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
			lifecycle:   life,
		}
	} else {
		err = fmt.Errorf("Type cannot be used as service: %v", factoryFuncType)

	}
	return
}

var contextReference = (*context.Context)(nil)
var contextReferenceType = reflect.TypeOf(contextReference).Elem()

// The resolveServiceFromValue function is called to resolve a service,
// and its arguments are a Context and a Value that is a pointer to a variable whose type is the interface to be resolved
// (this will make more sense when you see a service resolution in action).
// To resolve a service, the resolveServiceFromValue function looks to see if there is a BindingMap in the services map,
// using the requested type as the key.
// If there is a BindingMap, then its factory function is invoked, and the value is assigned via the pointer.
func resolveServiceFromValue(c context.Context, val reflect.Value) (err error) {
	serviceType := val.Elem().Type()
	if serviceType == contextReferenceType {
		val.Elem().Set(reflect.ValueOf(c))
	} else if binding, found := services[serviceType]; found {
		if binding.lifecycle == Scoped {
			resolveScopedService(c, val, binding)
		} else {
			val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
		}
	} else {
		err = fmt.Errorf("Cannot find service %v", serviceType)
	}
	return
}

// Special handling is required for scoped services.
// The resolveScopedService checks to see if the Context contains a value from a previous request to resolve the service.
// If not, the service is resolved and added to the Context so that it can be reused within the same scope.
func resolveScopedService(c context.Context, val reflect.Value,
	binding BindingMap) (err error) {
	sMap, ok := c.Value(ServiceKey).(serviceMap)
	if ok {
		serviceVal, ok := sMap[val.Type()]
		if !ok {
			serviceVal = invokeFunction(c, binding.factoryFunc)[0]
			sMap[val.Type()] = serviceVal
		}
		val.Elem().Set(serviceVal)
	} else {
		val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
	}
	return
}

func resolveFunctionArguments(c context.Context, f reflect.Value,
	otherArgs ...interface{}) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	if otherArgs != nil {
		for ; i < len(otherArgs); i++ {
			params[i] = reflect.ValueOf(otherArgs[i])
		}
	}
	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		err := resolveServiceFromValue(c, pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}

// The invokeFunction function is responsible for calling the factory function,
// using the resolveFunctionArguments function to inspect the factory function’s parameters and resolve each of them.
// These functions accept optional additional arguments, which are used when a function should be invoked with a mix of services
// and regular value parameters (in which case the parameters that require regular values must be defined first).
func invokeFunction(c context.Context, f reflect.Value,
	otherArgs ...interface{}) []reflect.Value {
	return f.Call(resolveFunctionArguments(c, f, otherArgs...))
}
