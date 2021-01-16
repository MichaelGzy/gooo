package server

import "github.com/MichaelGzy/gooo/go-micro/v1/auth"

type Options struct {
	Name string
	Auth auth.Auth
}

// Server name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}
