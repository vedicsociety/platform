/*
The User interface will represent an authenticated user so that requests to restricted resources can be evaluated.
*/
package identity

type User interface {
	GetID() int

	GetDisplayName() string

	InRole(name string) bool

	IsAuthenticated() bool

	IsPasswordTrue(pass string) bool
}
