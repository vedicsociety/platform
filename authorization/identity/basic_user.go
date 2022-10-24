/*
It is a default implementation of the User interface, which will be useful for applications with simple authorization requirements.
*/
package identity

import "strings"

// UnauthenticatedUser variable will be used to represent a user who has not signed into the application.
var UnauthenticatedUser User = &basicUser{}

// The NewBasicUser function creates a simple implementation of the User interface
func NewBasicUser(id int, name string, roles ...string) User {
	return &basicUser{
		Id:            id,
		Name:          name,
		Roles:         roles,
		Authenticated: true,
	}
}

type basicUser struct {
	Id            int
	Name          string
	Roles         []string
	Authenticated bool
}

func (user *basicUser) GetID() int {
	return user.Id
}

func (user *basicUser) GetDisplayName() string {
	return user.Name
}

func (user *basicUser) InRole(role string) bool {
	for _, r := range user.Roles {
		if strings.EqualFold(r, role) {
			return true
		}
	}
	return false
}

func (user *basicUser) IsAuthenticated() bool {
	return user.Authenticated
}

func (user *basicUser) IsPasswordTrue(pass string) bool {
	return true
}
