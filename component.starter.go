package main

import "html/template"

type Starter struct {
	Installed []StarterLink
	Resources []StarterLink
}

type StarterLink struct {
	Icon        template.HTML
	Title       string
	Description string
	Href        string
}
