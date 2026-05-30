package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Logger Request")
		start := time.Now()

		next.ServeHTTP(w, r) // পরবর্তী হ্যান্ডলারকে কল

		titme := time.Since(start) // হ্যান্ডলার সম্পন্ন হতে কত সময় লেগেছে তা হিসাব করা

		log.Println("Request Method ", r.Method, ", Request URL ", r.URL.Path, ", Duration ", titme) // লগ করা

	})

}

// http.Handler এক্সপেক্ট করে > একটি (htttp.HandlerFunc, http.HandlerFunc) এক্সেপক্ট করে > একটি ফাংশন যা http.ResponseWriter এবং *http.Request প্যারামিটার নেয়।
