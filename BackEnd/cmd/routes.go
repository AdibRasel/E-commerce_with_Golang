package cmd

import (
	"BackEnd/handlers"
	"BackEnd/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /products",
		manager.With( // সকল Middlware কে কল করা হচ্ছে
			http.HandlerFunc(handlers.GetProducts), // GetProducts কে কল করা হচ্ছে।
		),
	)

	mux.Handle(
		"POST /products",
		manager.With( // সকল Middlware কে কল করা হচ্ছে
			http.HandlerFunc(handlers.CreateProduct), // CreateProduct কে কল করা হচ্ছে।
		),
	)

	mux.Handle(
		"GET /products/{productId}",
		manager.With( // সকল Middlware কে কল করা হচ্ছে
			http.HandlerFunc(handlers.GetProductByID), // CreateProduct কে কল করা হচ্ছে।
		),
	)

}
