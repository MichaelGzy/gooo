package plugin

var (
	tmpl = `
package main

import (
	"github.com/MichaelGzy/gooo/go-micro/v0/micro/go-micro/v2/plugin"

	"{{.Path}}"
)

var Plugin = plugin.Config{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
