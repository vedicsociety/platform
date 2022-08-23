// To provide a service that presents the signed-in user

package authorization

import (
	"github.com/vedicsociety/platform/authorization/identity"
	"github.com/vedicsociety/platform/services"
	"github.com/vedicsociety/platform/sessions"
)

// The RegisterDefaultUserService function creates a scoped service for the User interface,
// which reads the value stored in the current session and uses it to query the UserStore service.
func RegisterDefaultUserService() {
	err := services.AddScoped(func(session sessions.Session,
		store identity.UserStore) identity.User {
		userID, found := session.GetValue(USER_SESSION_KEY).(int)
		if found {
			user, userFound := store.GetUserByID(userID)
			if userFound {
				return user
			}
		}
		return identity.UnauthenticatedUser
	})
	if err != nil {
		panic(err)
	}
}
