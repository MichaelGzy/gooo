package main

import (
	"fmt"
	"github.com/MichaelGzy/gooo/go-micro/v1"
	"github.com/MichaelGzy/gooo/oss/v2"
	"log"
	"runtime/debug"
)

func init() {
	fmt.Println("welcome to gooo")
}

func main() {
	fmt.Println("hello world")
	oss.SayHi()
	var options []micro.Option
	options = append(options, micro.Name("aaa"))

	service := micro.NewService(options...)
	service.Init()
	err := service.Run()
	if err != nil {
		log.Println(err)
		err := recover()
		if err != nil {
			log.Println(err)
			debug.PrintStack()
		}

	}
}
