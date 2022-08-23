// Create basic implementations of the authorization features that an application using the platform will provide.

package placeholder

import (
	"strings"

	"github.com/vedicsociety/platform/authorization/identity"
	"github.com/vedicsociety/platform/services"
)

func RegisterPlaceholderUserStore() {
	err := services.AddSingleton(func() identity.UserStore {
		return &PlaceholderUserStore{}
	})
	if err != nil {
		panic(err)
	}
}

var users = map[int]identity.User{
	1: identity.NewBasicUser(1, "Alice", "Administrator"),
	2: identity.NewBasicUser(2, "Bob"),
}

// The PlaceholderUserStore struct implements the UserStore interface with statically defined data for two users, Alice and Bob,
// and is used by the RegisterPlaceholderUserStore function to create a singleton service.
type PlaceholderUserStore struct{}

func (store *PlaceholderUserStore) GetUserByID(id int) (identity.User, bool) {
	user, found := users[id]
	return user, found
}

func (store *PlaceholderUserStore) GetUserByName(name string) (identity.User, bool) {
	for _, user := range users {
		if strings.EqualFold(user.GetDisplayName(), name) {
			return user, true
		}
	}
	return nil, false
}
