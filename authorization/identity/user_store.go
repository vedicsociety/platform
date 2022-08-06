package identity

// The user store provides access to the users who are known to the application, which can be located by ID or name.
type UserStore interface {
	GetUserByID(id int) (user User, found bool)

	GetUserByName(name string) (user User, found bool)
}
