# GRPC Client

The grpc client is a [micro.Client](https://godoc.org/github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/client#Client) compatible client.

## Overview

The client makes use of the [google.golang.org/grpc](google.golang.org/grpc) framework for the underlying communication mechanism.

## Usage

Specify the client to your micro service

```go
import (
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro"
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-plugins/client/grpc"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Client(grpc.NewClient()),
	)
}
```
