package etcd

import (
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/registry"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
