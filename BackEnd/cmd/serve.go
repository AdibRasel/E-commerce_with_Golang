package cmd

import (
	"BackEnd/middleware"
	"fmt"
	"net/http"
)

func Serve() {


	// নতুন Middleware Manager তৈরি করা হচ্ছে
	// এটি global middleware গুলো পরিচালনা করবে
	manager := middleware.NewManager()
	manager.Use( // এখানে সবার শেষের টা আগে কল হবে, পিচন থেকে সামনে কল হবে।
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)


	mux := http.NewServeMux() // Router রান করার জন্য
	wrappedMux := manager.WrapMux(mux) // সকল Middlware Wrap করা হলো
	
	
	initRoutes(mux, manager) // সকল Routes এবং Middlware এখানে যাবে।

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", wrappedMux) // Start the server

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
