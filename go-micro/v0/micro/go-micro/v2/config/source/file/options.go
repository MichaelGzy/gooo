package file

import (
	"context"

	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/source"
)

type filePathKey struct{}

// WithPath sets the path to file
func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, filePathKey{}, p)
	}
}
