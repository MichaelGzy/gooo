package service

import (
	"context"

	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/client"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/registry"
)

type clientKey struct{}

// WithClient sets the RPC client
func WithClient(c client.Client) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, clientKey{}, c)
	}
}
