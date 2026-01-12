package types

type App struct {
	Entity
	Name   string    `json:"name"`
	Config AppConfig `json:"config"`
}

type AppConfig struct {
	// A Group's primary key. If RequireExplicitGroup is off, User objects get this one assigned.
	DefaultGroupID int `json:"default_group_id"`
	// RequireExplicitGroup controls whether a User object must have a specified group to
	// be created.
	RequireExplicitGroup bool `json:"require_explicit_group"`
	// AllowRegistrationForApp controls whether User accounts can be registered or not.
	AllowRegistrationForApp bool `json:"allow_registration_for_app"`
	// AllowLoginForApp will temporarily block all login and token validation attempts.
	AllowLoginForApp bool `json:"allow_login_for_app"`
	// RequireGroupInsider indicates whether anyone can join any group or if someone from within the group
	// must provision their account first.
	RequireGroupInsider bool `json:"require_group_insider"`
	// AllowCrossAppReads controls whether members of an App can read from other apps. This should *probably*
	// only be on for the Saturn Microservice app.
	AllowCrossAppReads bool `json:"allow_cross_app_reads"`
	// AllowCrossGroupReads controls whether members of a Group can read from other groups. This should *probably*
	// be disabled.
	AllowCrossGroupReads bool `json:"allow_cross_group_reads"`
}
