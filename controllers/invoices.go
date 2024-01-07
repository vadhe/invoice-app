package controllers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jackc/pgx/v4"
	"github.com/vadhe/invoice-app/models"
	view "github.com/vadhe/invoice-app/views"
)

func Invoices(db *pgx.Conn) {
	var invoices []models.Invoice

	invoices = models.GetInvoice(db)
	layout := view.Layout(view.Filter(len(invoices)), view.List(invoices, "Create an invoice by clicking the ", "New Invoice button and get started"))
	http.Handle("/", templ.Handler(layout))

	http.HandleFunc("/invoice", func(w http.ResponseWriter, r *http.Request) {
		draft := r.URL.Query().Get("draft")
		pending := r.URL.Query().Get("pending")
		paid := r.URL.Query().Get("paid")
		if draft == "draft" || pending == "pending" || paid == "paid" {
			invoices = models.GetInvoiceByStatus(db, draft, paid, pending)
		} else {
			invoices = models.GetInvoice(db)
		}
		list := view.List(invoices, "We cannot find your Filter", "You can try to change to other Status")
		list.Render(r.Context(), w)
	})

	http.HandleFunc("/invoice/detail/", func(w http.ResponseWriter, r *http.Request) {
		layout := view.Layout(view.Back(), view.Detail())
		layout.Render(r.Context(), w)
	})
}
