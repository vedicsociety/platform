package identity

// The SignInManager interface will be used to define a service that the application will use to sign a user into and out of the application.
// The details of how the user is authenticated are left to the application.
type SignInManager interface {
	SignIn(user User) error
	SignOut(user User) error
}
