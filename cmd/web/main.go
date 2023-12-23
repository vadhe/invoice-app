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
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
