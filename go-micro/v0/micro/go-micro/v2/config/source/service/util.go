package service

import (
	"time"

	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/source"
	proto "github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/config/source/service/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      []byte(c.Data),
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
