// Create a simple access condition that checks to see if a user is in a role.

package authorization

import (
	"github.com/tsiparinda/platform/authorization/identity"
)

// The NewRoleCondition function accepts a set of roles, which are used to create a condition
// that will return true if a user has been assigned to any one of them.
func NewRoleCondition(roles ...string) identity.AuthorizationCondition {
	return &roleCondition{allowedRoles: roles}
}

type roleCondition struct {
	allowedRoles []string
}

// check Is User in Role?
func (c *roleCondition) Validate(user identity.User) bool {
	for _, allowedRole := range c.allowedRoles {
		if user.InRole(allowedRole) {
			return true
		}
	}
	return false
}
