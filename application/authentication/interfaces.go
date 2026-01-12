package authentication

import (
	"context"
	"github.com/ray-laboratories/saturn/types"
)

// Service represents the authentication service interface that both
// the SDK and local versions will meet.
type Service interface {
	// Login attempts to authenticate a user by a username-password combo.
	// On success, it stores the session and provides you the token.
	Login(ctx context.Context, username, password string) (string, error)
	// Logout will take your token and delete your session.
	Logout(ctx context.Context, token string) error
	// Register takes a well-formed types.UserRequest and saves it. Then,
	// it runs Login.
	Register(ctx context.Context, req *types.UserRequest) (string, error)
	// Validate takes a token and returns the types.User that is currently
	// logged in.
	Validate(ctx context.Context, token string) (*types.User, error)
}
