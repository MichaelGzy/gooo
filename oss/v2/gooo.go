package oss

import (
	"fmt"
	"github.com/MichaelGzy/gooo/srv/v2"
)

func SayHi() {
	fmt.Println("Hi_oss_v2")
	srv.SayHi()
}

func SayHello() {
	fmt.Println("Hello_oss_v2")
}
