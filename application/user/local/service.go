package local

import (
	"errors"
	"fmt"
	"github.com/ray-laboratories/saturn/application/user"
	"github.com/ray-laboratories/saturn/types"
)

var _ user.Service = (*UserService)(nil)

type UserService struct {
	userRepository  UserRepository
	groupRepository GroupRepository
}

func (u *UserService) checkPermissionOnUser(ac types.AuthenticatedContext, foundUser *types.User) (bool, error) {
	// Then, we make sure and double check two important things.
	// First, is the user part of the same group as us, and are we
	// authorized to read that?
	if foundUser.GroupID != ac.Group.ID && !ac.App.Config.AllowCrossGroupReads {
		return false, nil
	}

	// Then, we actually find that group.
	group, err := u.groupRepository.Get(ac.Context, foundUser.GroupID)
	if err != nil {
		return false, fmt.Errorf("getting group by id: %w", err)
	}

	// Now, we check our second thing: is the user part of the same
	// app-space as us, and are we allowed to read that?
	if group.AppID != ac.App.ID && !ac.App.Config.AllowCrossAppReads {
		return false, nil
	}
	return true, nil
}

func (u *UserService) GetUserByID(ac types.AuthenticatedContext, id uint) (*types.User, error) {
	// First, we find the user.
	userToFind, err := u.userRepository.Get(ac.Context, id)
	if err != nil {
		return nil, fmt.Errorf("getting userToFind by id: %w", err)
	}

	// Second, we perm-check.
	canRead, err := u.checkPermissionOnUser(ac, userToFind)
	if err != nil {
		return nil, fmt.Errorf("checking permission on user: %w", err)
	}
	if !canRead {
		return nil, errors.New("missing permissions on user")
	}

	// If all is OK, then return the user.
	return userToFind, nil
}

func (u *UserService) GetUserByUsername(ac types.AuthenticatedContext, appName, username string) (*types.User, error) {
	userToFind, err := u.userRepository.GetByUsername(ac.Context, appName, username)
	if err != nil {
		return nil, fmt.Errorf("getting user by username: %w", err)
	}
	canRead, err := u.checkPermissionOnUser(ac, userToFind)
	if err != nil {
		return nil, fmt.Errorf("checking permission on user: %w", err)
	}
	if !canRead {
		return nil, errors.New("missing permissions on user")
	}
	return userToFind, nil
}

func (u *UserService) UpdateUser(ac types.AuthenticatedContext, user *types.User) (*types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(ac types.AuthenticatedContext, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetUsersInGroup(ac types.AuthenticatedContext, groupName string, limit, offset uint) ([]*types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetUsersInApp(ac types.AuthenticatedContext, appName string, limit, offset uint) ([]*types.User, error) {
	//TODO implement me
	panic("implement me")
}
