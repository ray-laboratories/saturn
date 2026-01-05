package types

type Session struct {
	Token    string
	Username string
}

func NewSession(id, username string) *Session {
	return &Session{
		Token:    id,
		Username: username,
	}
}
