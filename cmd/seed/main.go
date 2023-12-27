package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type Invoice struct {
	ID         string
	IdInvoices string
	Status     string
	Price      float64
	Name       string
}

var statuses = []string{"paid", "pending", "draft"}

func generateRandomIdInvoices() string {
	return "#RT" + strconv.Itoa(rand.Intn(5000)+1000)
}

func generateRandomStatus() string {
	return statuses[rand.Intn(len(statuses))]
}

func generateRandomPrice() float64 {
	return rand.Float64() * 2000
}

func generateRandomName() string {
	names := []string{"Jensen Huang", "John Doe", "Jane Doe", "Alice Smith", "Bob Johnson"}
	return names[rand.Intn(len(names))]
}

func main() {
	// Check env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Connect to DB
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	insertInvoice(db)
}

func insertInvoice(db *pgx.Conn) {
	for i := 0; i < 5; i++ {
		randomInvoice := Invoice{
			IdInvoices: generateRandomIdInvoices(),
			Status:     generateRandomStatus(),
			Price:      generateRandomPrice(),
			Name:       generateRandomName(),
		}

		_, err := db.Exec(context.Background(),
			"INSERT INTO invoices (id_invoices, status, price, name) VALUES ($1, $2, $3, $4)",
			randomInvoice.IdInvoices, randomInvoice.Status, randomInvoice.Price, randomInvoice.Name,
		)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Inserted invoice with ID %s\n", randomInvoice.ID)
	}
}
