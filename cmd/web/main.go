package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	view "github.com/vadhe/invoice-app/views"
)

func main() {
	// check env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// connect to DB
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	fmt.Println("Successfully connected to the database!")

	// render index page
	component := view.Index()
	http.Handle("/", templ.Handler(component))

	// make folder dist and public as static file
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
