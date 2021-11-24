package main

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
)

type PageIndex struct {
	Navbar  kyoto.Component
	Starter kyoto.Component
}

func (p *PageIndex) Template() *template.Template {
	return newtemplate("page.index.html")
}

func (p *PageIndex) Meta() kyoto.Meta {
	return kyoto.Meta{
		Title: "kyoto - Project Starter",
	}
}

func (p *PageIndex) Init() {
	p.Navbar = kyoto.RegC(p, UINavbar())
	p.Starter = kyoto.RegC(p, &Starter{
		Installed: []StarterLink{
			{
				Icon:        `<img src="/static/img/icons/kyoto.svg" class="h-8 w-8" />`,
				Title:       "Kyoto",
				Description: "Core library for building frontend",
				Href:        "https://github.com/kyoto-framework/kyoto",
			},
			{
				Icon:        `<img src="/static/img/icons/uikit.svg" class="ml-0.5 h-6 w-6" />`,
				Title:       "Kyoto UIKit",
				Description: "Set of components for rapid development",
				Href:        "https://github.com/kyoto-framework/uikit",
			},
			{
				Icon:        `<img src="/static/img/icons/tailwind.svg" class="h-8 w-8" />`,
				Title:       "Tailwind CSS",
				Description: "A utility-first CSS framework",
				Href:        "https://github.com/kyoto-framework/kyoto",
			},
		},
	})
}
