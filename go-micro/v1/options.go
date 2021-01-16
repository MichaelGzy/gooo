package micro

import (
	"context"
	"github.com/MichaelGzy/gooo/go-micro/v1/auth"
	"github.com/MichaelGzy/gooo/go-micro/v1/client"
	"github.com/MichaelGzy/gooo/go-micro/v1/debug/profile"
	"github.com/MichaelGzy/gooo/go-micro/v1/registry"
	"github.com/MichaelGzy/gooo/go-micro/v1/server"
)

type Options struct {
	Auth auth.Auth
	//Broker    broker.Broker
	//Cmd       cmd.Cmd
	//Config    config.Config
	Client client.Client
	Server server.Server
	//Store     store.Store
	//Registry  registry.Registry
	//Runtime   runtime.Runtime
	//Transport transport.Transport
	Profile profile.Profile

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Auth:      auth.DefaultAuth,
		Broker:    broker.DefaultBroker,
		Cmd:       cmd.DefaultCmd,
		Config:    config.DefaultConfig,
		Client:    client.DefaultClient,
		Server:    server.DefaultServer,
		Store:     store.DefaultStore,
		Registry:  registry.DefaultRegistry,
		Runtime:   runtime.DefaultRuntime,
		Transport: transport.DefaultTransport,
		Context:   context.Background(),
		Signal:    true,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func Name(n string) Option {
	return func(o *Options) {
		o.Server.Init(server.Name(n))
	}
}
