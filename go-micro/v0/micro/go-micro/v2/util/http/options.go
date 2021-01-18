package http

import (
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/registry"
)

type Options struct {
	Registry registry.Registry
}

type Option func(*Options)

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}
