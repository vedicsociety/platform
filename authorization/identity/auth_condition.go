package identity

// The AuthorizationCondition interface will be used to assess whether a signed-in user has access to a protected URL
// and will be used as part of the request handling process.
type AuthorizationCondition interface {
	Validate(user User) bool
}
