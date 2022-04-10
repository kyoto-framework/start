package main

import (
	"html/template"

	"github.com/kyoto-framework/i18n"
	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/render"
	"github.com/kyoto-framework/zen"
)

func FuncMap(c *kyoto.Core) template.FuncMap {
	return render.ComposeFuncMap(
		render.FuncMap(c), // Base kyoto functions
		zen.FuncMap(),     // Zen functions
		i18n.FuncMap(),    // i18n functions
	)
}

func Template(c *kyoto.Core, page string) *template.Template {
	tmpl := template.New(page)                                // Initialize template
	tmpl = tmpl.Funcs(FuncMap(c))                             // Apply funcmap
	tmpl = template.Must(tmpl.ParseGlob("*.html"))            // Parse base project files
	tmpl = template.Must(tmpl.ParseGlob("uikit/twui/*.html")) // Parse components library
	return tmpl
}
