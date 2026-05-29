package cmd

import (
	"BackEnd/global_router"
	"BackEnd/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() // Router



	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))


	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))





	globalRouter := global_router.GlobalRouter(mux) // globalRouter ফাংশনটি একটি middleware হিসেবে কাজ করবে,
	// যা সব রিকোয়েস্ট এর আগে CORS policy হ্যান্ডেল করবে এবং Preflight রিকোয়েস্ট হ্যান্ডেল করবে।

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", globalRouter) // Start the server

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
