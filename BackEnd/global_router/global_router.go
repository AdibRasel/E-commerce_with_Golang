package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {

	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                         // অন্য ডোমেইন থেকে এই API কে access করা যায় * দ্বারা।
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE") // এই API তে  GET, POST, PATCH, DELETE রিকোয়েস্ট আসবে।
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")             // এই API তে শুধু মাত্র Content-Type হেডার আসবে।
		w.Header().Set("Content-Type", "application/json")                         // json format hobe.

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)

}
