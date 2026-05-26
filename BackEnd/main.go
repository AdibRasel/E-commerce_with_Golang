package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func abountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

type Product struct {
	ID          int     `json:"id"`          // struct এর বিতরে প্রোপার্টির প্রথম অক্ষর বড় হাতের দিতে হবে। না হলে এই প্রোপার্টি json এ কনভার্ট হবে না।
	Title       string  `json:"title"`       // struct এর প্রথম অক্ষর ছোট হাতের দিলে সেটি প্রাইভেট হয়ে যায়, এটি শুধু মাত্র মেইন প্যাকেজের মধ্যে ব্যবহার করা যাবে, অন্য প্যাকেজে বাবহার করা যাবে না।
	Description string  `json:"description"` // struct এর প্রোপার্টির নাম json এ কনভার্ট করার সময় description হবে।
	Price       float64 `json:"price"`
	// ImgURL      string  `json:"ImgURL_er_poriborthe_ei_nam_dekhabe"` //
	ImgURL string `json:"imageUrl"` //

}

var productList []Product

// getProduct API
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS policy set করা হচ্ছে, যাতে অন্য ডোমেইন থেকে এই API কে access করা যায়।
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")  এখন শুধু মাত্র ৩০০০ পোর্ট এর ডোমেইন এক্সেস করতে পারবে।

	w.Header().Set("Content-Type", "application/json") // json format hobe.

	if r.Method != http.MethodGet { //r.method = post, put, patch, delete
		http.Error(
			w, "Please give me Get Request", // Response body
			400, // Bad Request status code
		)
		return
	}

	encoder := json.NewEncoder(w) // w = http.ResponseWriter // w মানি হচ্ছে রেসপন্স পাঠানো।

	encoder.Encode(productList) //struct -> ProductList convert json format.

}

// createProduct API
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             // CORS policy set করা হচ্ছে, যাতে অন্য ডোমেইন থেকে এই API কে access করা যায়।
	w.Header().Set("Access-Control-Allow-Methods", "POST")         // এই API তে শুধু মাত্র POST রিকোয়েস্ট আসবে।
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // এই API তে শুধু মাত্র Content-Type হেডার আসবে।
	w.Header().Set("Content-Type", "application/json")             // json format hobe.

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" { //r.method = get, post, put, patch, delete
		http.Error(
			w, "Please give me Post Request", // Response body
			400, // Bad Request status code
		)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)


	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400) // w, err.Error(), 400
		
		return
	}

	newProduct.ID = len(productList) + 1 // new product এর id set করা হচ্ছে, যাতে করে প্রতিটি প্রোডাক্ট এর আলাদা আলাদা id থাকে।
	productList = append(productList, newProduct)

	w.WriteHeader(201)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct) // নতুন প্রোডাক্ট এর ডাটা রেসপন্স হিসেবে পাঠানো হচ্ছে।


}


func main() {
	mux := http.NewServeMux() // Router

	mux.HandleFunc("/hello", helloHandler) // Route

	mux.HandleFunc("/about", abountHandler) // Route

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products", createProduct)

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux) // Start the server

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I love orange.",
		Price:       100,
		ImgURL:      "https://cdn.britannica.com/24/174524-050-A851D3F2/Oranges.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I love apple.",
		Price:       120,
		ImgURL:      "https://www.paperandtea.com/cdn/shop/articles/Apfel_7ebe153a-a4ac-473a-9217-658894dfc968.jpg?v=1765535477&width=1500",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banna",
		Description: "Banna is Lovely.. I love Banna.",
		Price:       80,
		ImgURL:      "https://blog-images-1.pharmeasy.in/blog/production/wp-content/uploads/2021/01/30152155/shutterstock_518328943-1.jpg",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "Mango is Lovely.. I love Mango.",
		Price:       160,
		ImgURL:      "https://www.paperandtea.com/cdn/shop/articles/Mango_6fb74c95-c9b0-4559-88e8-f542e6d6b18d.jpg?v=1769533193&width=1024",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Grape",
		Description: "Grape is purple. I love grape.",
		Price:       200,
		ImgURL:      "https://img.imageboss.me/fourwinds/width/425/dpr:2/shop/products/flamegrapes2.jpg?v=1729716550",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Pomegranate",
		Description: "Pomegranate is red. I love Pomegranate.",
		Price:       200,
		ImgURL:      "https://www.gettystewart.com/wp-content/uploads/2022/11/pomegranate-and-seeds-s.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)

}
