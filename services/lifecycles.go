// Services are registered with lifecycles, which specify when the factory function is invoked to create new struct values.
// We use three service lifecycles
// Transient - For this lifecycle, the factory function is invoked for every service request.
// Singleton - For this lifecycle, the factory function is invoked once, and every request receives the same struct instance.
// Scoped - For this lifecycle, the factory function is invoked once for the first request within a scope,
// and every request within that scope receives the same struct instance.
package services

type lifecycle int

// 0,1,2
const (
	Transient lifecycle = iota
	Singleton
	Scoped
)
