package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/kyoto-framework/i18n"
	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/actions"
	"github.com/kyoto-framework/kyoto/render"
	"github.com/kyoto-framework/kyoto/smode"
	"github.com/kyoto-framework/uikit/twui"
)

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", render.PageHandler(PageIndex))
}

func setupActions(mux *http.ServeMux) {
	// Register handler
	mux.HandleFunc("/internal/actions/", actions.Handler(func(c *kyoto.Core) *template.Template {
		return Template(c, "Actions")
	}))
	// Register dynamic components
	smode.Register(
		&twui.AppUINavNavbar{},
	)
}

func setupStatic(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/dist"))))
}

func setupI18N(mux *http.ServeMux) {
	i18n.Parse("./static/i18n")
}

func main() {
	// Initialize mux
	mux := http.NewServeMux()

	// Setup server
	setupStatic(mux)
	setupI18N(mux)
	setupActions(mux)
	setupRoutes(mux)

	// Run server
	if os.Getenv("PORT") == "" {
		log.Println("Listening on http://localhost:25025")
		http.ListenAndServe("localhost:25025", mux)
	} else {
		log.Println("Listening on http://0.0.0.0:" + os.Getenv("PORT"))
		http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	}
}
