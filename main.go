package main

import (
	"github.com/vedicsociety/platform/placeholder"
	"github.com/vedicsociety/platform/services"
)

func main() {
	// services is a dependency injection (DI), in which code that depends on an interface can obtain an implementation
	// without needing to select an underlying type or create an instance directly.
	// During application startup, the interfaces defined by the application will be added to a register,
	// along with a factory function that creates instances of an implementation struct.
	// So, for example, the platform.logger.Logger interface will be registered with a factory function that invokes the NewDefaultLogger function.
	// When an interface is added to the register, it is known as a service.
	// During execution, application components that need the features described by the service go to the registry
	// and request the interface they want.
	// The registry invokes the factory function and returns the struct that is created,
	// which allows the application component to use the interface features without knowing or specifying
	// which implementation struct will be used or how it is created.
	// Don’t worry if this doesn’t make sense—this can be a difficult topic to understand, and it becomes easier once you see it in action.
	services.RegisterDefaultServices()
	placeholder.Start()
}
