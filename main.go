package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/yuriizinets/kyoto"
)

func ssatemplate(p kyoto.Page) *template.Template {
	return mktemplate("SSA")
}

func ssahandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		kyoto.SSAHandlerFactory(ssatemplate, map[string]interface{}{
			"internal:rw": rw,
			"internal:r":  r,
		})(rw, r)
	}
}

func main() {
	// Routes
	http.HandleFunc("/", kyoto.PageHandler(&PageIndex{}))

	// Statics
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/dist"))))
	// SSA
	http.HandleFunc("/SSA/", ssahandler())

	// Run
	if os.Getenv("PORT") == "" {
		log.Println("Listening on http://localhost:25025")
		http.ListenAndServe("localhost:25025", nil)
	} else {
		log.Println("Listening on http://0.0.0.0:" + os.Getenv("PORT"))
		http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}
}
