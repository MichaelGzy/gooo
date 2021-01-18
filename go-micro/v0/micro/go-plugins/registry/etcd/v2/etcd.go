// Package etcd provides an etcd v3 service registry
package etcd

import (
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/cmd"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/registry"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/registry/etcd"
)

func init() {
	cmd.DefaultRegistries["etcd"] = etcd.NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return etcd.NewRegistry(opts...)
}
