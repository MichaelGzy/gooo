package provider

type Provider interface {
	String() string

	Options() Options

	EndPoint(...EndPointOption) string

	Redirect() string
}

type EndPointOptions struct {
	State string

	LoginHint string
}

type EndPointOption func(*EndPointOptions)

func WithState(c string) EndPointOption {
	return func(o *EndPointOptions) {
		o.State = c
	}
}

func WithHint(hint string) EndPointOption {
	return func(o *EndPointOptions) {
		o.LoginHint = hint
	}
}
