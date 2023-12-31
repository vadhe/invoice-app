package controllers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jackc/pgx/v4"
	"github.com/vadhe/invoice-app/models"
	view "github.com/vadhe/invoice-app/views"
)

func Invoices(db *pgx.Conn) {
	var invoices []models.Invoice

	invoices = models.GetInvoice(db)
	component := view.Index(invoices)
	http.Handle("/", templ.Handler(component))

	http.HandleFunc("/invoice", func(w http.ResponseWriter, r *http.Request) {
		draft := r.URL.Query().Get("draft")
		pending := r.URL.Query().Get("pending")
		paid := r.URL.Query().Get("paid")
		if draft == "draft" || pending == "pending" || paid == "paid" {
			invoices = models.GetInvoiceByStatus(db, draft, paid, pending)
		} else {
			invoices = models.GetInvoice(db)
		}
		// if len(invoices) == 0 {

		// }
		fmt.Println(len(models.GetInvoiceByStatus(db, draft, paid, pending)))
		component := view.List(invoices)
		component.Render(r.Context(), w)
	})
}
