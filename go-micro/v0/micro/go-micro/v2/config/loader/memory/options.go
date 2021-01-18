package memory

import (
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/loader"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/reader"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
