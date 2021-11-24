package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/kyoto-framework/kyoto"
)

func ssatemplate(p kyoto.Page) *template.Template {
	return newtemplate("SSA")
}

func main() {
	// Routes
	http.HandleFunc("/", kyoto.PageHandler(&PageIndex{}))

	// Statics
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/dist"))))
	// SSA
	http.HandleFunc("/SSA/", kyoto.SSAHandler(ssatemplate))

	// Run
	if os.Getenv("PORT") == "" {
		log.Println("Listening on http://localhost:25025")
		http.ListenAndServe("localhost:25025", nil)
	} else {
		log.Println("Listening on http://0.0.0.0:" + os.Getenv("PORT"))
		http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}
}
