package types

type User struct {
	Entity
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	GroupID        uint   `json:"group_id"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	GroupID  uint   `json:"group_id"`
}

func NewUser(username, hashedPassword string, groupID uint) *User {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		GroupID:        groupID,
	}
}
