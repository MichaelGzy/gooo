package store

import (
	"context"
	"github.com/MichaelGzy/gooo/go-micro/v1/client"
)

type Options struct {
	Nodes []string

	DataBase string

	Table string

	Context context.Context

	Client client.Client
}
