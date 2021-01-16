package server

import "github.com/MichaelGzy/gooo/go-micro/v1/registry"

type Server interface {
	// Initialise options
	Init(...Option) error
	// Retrieve the options
	Options() Options

	Name() string

	Handle(Handler) error

	NewHandler(interface{}, ...HandlerOption) Handler
}

type Option func(*Options)

type Handler interface {
	Name() string
	Handler() interface{}
	Endpoints() []*registry.Endpoint
	Options() HandlerOptions
}

func InternalHandler(b bool) HandlerOption {
	return func(o *HandlerOptions) {
		o.Internal = b
	}
}
