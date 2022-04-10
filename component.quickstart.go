package main

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/lifecycle"
)

type ComponentQuickstartConfig struct {
	Installed []ComponentQuickstartConfigLink
}

type ComponentQuickstartConfigLink struct {
	Icon  template.HTML
	Text  string
	Href  string
	Title string
}

func ComponentQuickstart(conf *ComponentQuickstartConfig) func(*kyoto.Core) {
	return func(c *kyoto.Core) {
		lifecycle.Init(c, func() {
			c.State.Set("Conf", conf)
		})
	}
}
