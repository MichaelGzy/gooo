package auth

import (
	"github.com/MichaelGzy/gooo/go-micro/v1/auth/provider"
	"github.com/MichaelGzy/gooo/go-micro/v1/client"
	"github.com/MichaelGzy/gooo/go-micro/v1/store"
)

type Options struct {
	NameSpace string

	ID string

	Secret string

	Token *Token

	PublicKey string

	PrivateKey string

	Provider provider.Provider

	LoginUrl string

	Store store.Store

	Client client.Client

	Addrs []string
}
