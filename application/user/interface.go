package user

import (
	"github.com/ray-laboratories/saturn/types"
)

// The Service interface defines common actions related to the types.User.
// Here, you can get users by their unique ID, their username (scoped to app name),
// and a myriad of other helpful operations.
type Service interface {
	GetUserByID(ac types.AuthenticatedContext, id uint) (*types.User, error)
	GetUserByUsername(ac types.AuthenticatedContext, appName, username string) (*types.User, error)
	UpdateUser(ac types.AuthenticatedContext, user *types.User) (*types.User, error)
	DeleteUser(ac types.AuthenticatedContext, id int) error
	GetUsersInGroup(ac types.AuthenticatedContext, groupName string, limit, offset uint) ([]*types.User, error)
	GetUsersInApp(ac types.AuthenticatedContext, appName string, limit, offset uint) ([]*types.User, error)
}
