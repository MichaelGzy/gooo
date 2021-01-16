package auth

import "time"

type Auth interface {
	//Init the auth
	Init(opts ...Option)
	//Option set for auth
	Option() Options
}

type Option func(*Options)

type Token struct {
	AccessToken string `json:"access_token"`

	RefreshToken string `json:"refresh_token"`

	Created time.Time `json:"created"`

	Expiry time.Time `json:"expiry"`
}
