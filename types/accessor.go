package types

type Accessor struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

type AccessorRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAccessor(username, hashedPassword string) *Accessor {
	return &Accessor{
		Username:       username,
		HashedPassword: hashedPassword,
	}
}
