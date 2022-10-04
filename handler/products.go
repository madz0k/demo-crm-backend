package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func products(router chi.Router) {
	router.Get("/", getAllProducts)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := dbInstance.GetAllProducts()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, products); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
