// To allow some simple authentication
// This request handler has a hardwired password— mysecret —for all users.
// The GetSignIn method displays a template to collect the user’s name and password.
// The PostSignIn method checks the password and makes sure there is a user in the store with the specified name,
// before signing the user into the application.
// The PostSignOut method signs the user out of the application.
// To create the template used by the handler, add a file named signin.html to the placeholder folder.
package placeholder

import (
	"fmt"

	"github.com/vedicsociety/platform/authorization/identity"
	"github.com/vedicsociety/platform/http/actionresults"
)

type AuthenticationHandler struct {
	identity.User
	identity.SignInManager
	identity.UserStore
}

func (h AuthenticationHandler) GetSignIn() actionresults.ActionResult {
	return actionresults.NewTemplateAction("signin.html",
		fmt.Sprintf("Signed in as: %v", h.User.GetDisplayName()))
}

type Credentials struct {
	Username string
	Password string
}

func (h AuthenticationHandler) PostSignIn(creds Credentials) actionresults.ActionResult {
	if creds.Password == "mysecret" {
		user, ok := h.UserStore.GetUserByName(creds.Username)
		if ok {
			h.SignInManager.SignIn(user)
			return actionresults.NewTemplateAction("signin.html",
				fmt.Sprintf("Signed in as: %v", user.GetDisplayName()))
		}
	}
	return actionresults.NewTemplateAction("signin.html", "Access Denied")
}

func (h AuthenticationHandler) PostSignOut() actionresults.ActionResult {
	h.SignInManager.SignOut(h.User)
	return actionresults.NewTemplateAction("signin.html", "Signed out")
}
