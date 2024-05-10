package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/health", app.Health)

	mux.Post("/products", app.CreateProduct)
	mux.Get("/products/get/{id}", app.GetProduct)
	mux.Put("/products/update/{id}", app.UpdateProduct)
	mux.Get("/products/all", app.AllProducts)
	mux.Delete("/products/delete/{id}", app.DeleteProduct)
	return mux
}
