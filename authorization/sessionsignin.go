// Implement the interfaces that the platform will provide for authorization.
package authorization

import (
	"context"
	"platform/authorization/identity"
	"platform/services"
	"platform/sessions"
)

const USER_SESSION_KEY string = "USER"

// The RegisterDefaultSignInService function creates a scoped service for the SignInManager interface,
// which is resolved using the SessionSignInMgr struct.
func RegisterDefaultSignInService() {
	err := services.AddScoped(func(c context.Context) identity.SignInManager {
		return &SessionSignInMgr{Context: c}
	})
	if err != nil {
		panic(err)
	}
}

// The SessionSignInMgr struct implements the SignInManager interface by storing the signed-in user’s ID in the session
// and removing it when the user is signed out.
// Relying on sessions ensures that a user will remain signed in until they sign out or their session expires.
type SessionSignInMgr struct {
	context.Context
}

func (mgr *SessionSignInMgr) SignIn(user identity.User) (err error) {
	session, err := mgr.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, user.GetID())
	}
	return
}

func (mgr *SessionSignInMgr) SignOut(user identity.User) (err error) {
	session, err := mgr.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, nil)
	}
	return
}

func (mgr *SessionSignInMgr) getSession() (s sessions.Session, err error) {
	err = services.GetServiceForContext(mgr.Context, &s)
	return
}
