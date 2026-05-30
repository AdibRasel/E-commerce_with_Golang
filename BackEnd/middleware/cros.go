package middleware

import (
	"fmt"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Cros Request")
		// handle cors
		w.Header().Set("Access-Control-Allow-Origin", "*")                         // অন্য ডোমেইন থেকে এই API কে access করা যায় * দ্বারা।
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE") // এই API তে  GET, POST, PATCH, DELETE রিকোয়েস্ট আসবে।
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")             // এই API তে শুধু মাত্র Content-Type হেডার আসবে।
		w.Header().Set("Content-Type", "application/json")                         // json format hobe.
		next.ServeHTTP(w, r)
	})
}
