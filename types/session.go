package types

import "context"

type Session struct {
	Token string
	User  *User
	Group *Group
	App   *App
}

func NewSession(token string, user *User) *Session {
	return &Session{
		Token: token,
		User:  user,
	}
}

type AuthenticatedContext struct {
	context.Context
	*Session
}
