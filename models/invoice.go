package models

import (
	"context"
	"log"

	"time"

	"github.com/jackc/pgx/v4"
	"github.com/vadhe/invoice-app/libs"
)

type Invoice struct {
	ID         int
	IdInvoices string
	Status     string
	Date       time.Time
	Price      float64
	Name       string
}

func InsertInvoice(db *pgx.Conn) {
	for i := 0; i < 5; i++ {
		randomInvoice := Invoice{
			IdInvoices: libs.GenerateRandomIdInvoices(),
			Status:     libs.GenerateRandomStatus(),
			Price:      libs.GenerateRandomPrice(),
			Name:       libs.GenerateRandomName(),
		}

		_, err := db.Exec(context.Background(),
			"INSERT INTO invoices (id_invoices, status, price, name) VALUES ($1, $2, $3, $4)",
			randomInvoice.IdInvoices, randomInvoice.Status, randomInvoice.Price, randomInvoice.Name,
		)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetInvoice(db *pgx.Conn) []Invoice {
	rows, err := db.Query(context.Background(), "SELECT * FROM invoices")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var invoices []Invoice

	for rows.Next() {
		var invoice Invoice
		err := rows.Scan(&invoice.ID, &invoice.Date, &invoice.IdInvoices, &invoice.Status, &invoice.Price, &invoice.Name)
		if err != nil {
			log.Fatal(err)
		}
		invoices = append(invoices, invoice)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return invoices
}
