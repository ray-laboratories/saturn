package types

type Group struct {
	Entity
	Name  string `json:"name"`
	AppID uint   `json:"app_id"`
}
