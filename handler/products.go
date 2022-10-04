package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/madz0k/demo-crm-backend/db"
	"github.com/madz0k/demo-crm-backend/models"
	"net/http"
	"strconv"
)

var productIDKey = "productID"

func products(router chi.Router) {
	router.Get("/", getAllProducts)
	router.Post("/", createProduct)

	router.Route("/{productId}", func(router chi.Router) {
		router.Use(ProductContext)
		router.Get("/", getProduct)
		router.Put("/", updateProduct)
		router.Delete("/", deleteProduct)
	})
}

func ProductContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "productId")
		if productId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("product ID is required")))
			return
		}
		id, err := strconv.Atoi(productId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid product ID")))
		}
		ctx := context.WithValue(r.Context(), productIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

func createProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	if err := render.Bind(r, product); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddProduct(product); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, product); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.Context().Value(productIDKey).(int)
	product, err := dbInstance.GetProductById(productID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &product); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.Context().Value(productIDKey).(int)
	err := dbInstance.DeleteProduct(productId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.Context().Value(productIDKey).(int)
	productData := models.Product{}
	if err := render.Bind(r, &productData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	product, err := dbInstance.UpdateProduct(productId, productData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &product); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
