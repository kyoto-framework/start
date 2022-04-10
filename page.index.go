package main

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/lifecycle"
	"github.com/kyoto-framework/kyoto/render"
)

func PageIndex(c *kyoto.Core) {
	render.Template(c, func() *template.Template {
		return Template(c, "page.index.html")
	})
	lifecycle.Init(c, func() {
		c.Component("Navbar", ComponentNavbar())
		c.Component("Quickstart", ComponentQuickstart(&ComponentQuickstartConfig{
			Installed: []ComponentQuickstartConfigLink{
				{
					Icon:  `<img src="/static/img/icons/kyoto.svg" class="h-8 w-8" />`,
					Title: "Kyoto",
					Text:  "Core library for building frontend",
					Href:  "https://github.com/kyoto-framework/kyoto",
				},
				{
					Icon:  `<img src="/static/img/icons/kyoto.svg" class="h-8 w-8" />`,
					Title: "Zen",
					Text:  "Set of commonly used Go functions",
					Href:  "https://github.com/kyoto-framework/kyoto",
				},
				{
					Icon:  `<img src="/static/img/icons/kyoto.svg" class="h-8 w-8" />`,
					Title: "i18n",
					Text:  "Internationalizing helper",
					Href:  "https://github.com/kyoto-framework/kyoto",
				},
				{
					Icon:  `<img src="/static/img/icons/uikit.svg" class="ml-0.5 h-6 w-6" />`,
					Title: "Kyoto UIKit",
					Text:  "Set of components for rapid development",
					Href:  "https://github.com/kyoto-framework/uikit",
				},
				{
					Icon:  `<img src="/static/img/icons/tailwind.svg" class="h-8 w-8" />`,
					Title: "Tailwind CSS",
					Text:  "An utility-first CSS framework",
					Href:  "https://github.com/kyoto-framework/kyoto",
				},
			},
		}))
	})
}
