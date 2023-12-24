package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	view "github.com/vadhe/invoice-app/views"
)

func main() {
	component := view.Index()

	http.Handle("/", templ.Handler(component))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
