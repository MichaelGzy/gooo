package micro

import (
	"github.com/MichaelGzy/gooo/go-micro/v1/auth"
	"github.com/MichaelGzy/gooo/go-micro/v1/client"
	"github.com/MichaelGzy/gooo/go-micro/v1/debug/trace"
	"github.com/MichaelGzy/gooo/go-micro/v1/server"
	"github.com/MichaelGzy/gooo/go-micro/v1/util/wrapper"
)

type Service interface {
	Name() string
	Init(...Option)
	Option() Options
	Client() client.Client
	Server() server.Server
	Run() error
	String() string
}

type Option func(*Options)

// NewService creates and returns a new Service based on the packages within.
func NewService(opts ...Option) Service {
	return newService(opts...)
}

func newService(opts ...Option) Service {
	service := new(service)
	options := newOptions(opts...)

	// service name
	serviceName := options.Server.Options().Name

	// we pass functions to the wrappers since the values can change during initialisation
	authFn := func() auth.Auth { return options.Server.Options().Auth }
	cacheFn := func() *client.Cache { return options.Client.Options().Cache }

	// wrap client to inject From-Service header on any calls
	options.Client = wrapper.FromService(serviceName, options.Client)
	options.Client = wrapper.TraceCall(serviceName, trace.DefaultTracer, options.Client)
	options.Client = wrapper.CacheClient(cacheFn, options.Client)
	options.Client = wrapper.AuthClient(authFn, options.Client)

	// wrap the server to provide handler stats
	options.Server.Init(
	//server.WrapHandler(wrapper.HandlerStats(stats.DefaultStats)),
	//server.WrapHandler(wrapper.TraceHandler(trace.DefaultTracer)),
	//server.WrapHandler(wrapper.AuthHandler(authFn)),
	)

	// set opts
	service.opts = options

	return service
}
