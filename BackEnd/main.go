package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func abountHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}


type Product struct {
	ID int
	Title string 
	Description string 
	Price float64 
	ImgURL string
}

var productList []Product 


func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{ //r.method = post, put, patch, delete
		http.Error(
			w, "Please give me Get Request", // Response body
			400, // Bad Request status code 
		)
		return
	}

	encoder := json.NewEncoder(w) // w = http.ResponseWriter

	encoder.Encode(productList)

}


func main(){
	mux := http.NewServeMux() // Router

	mux.HandleFunc("/hello", helloHandler) // Route

	mux.HandleFunc("/about", abountHandler) // Route


	mux.HandleFunc("/products", getProducts)




	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux) // Start the server

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}



func init(){
	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Orange is red. I love orange.",
		Price: 100,
		ImgURL: "https://cdn.britannica.com/24/174524-050-A851D3F2/Oranges.jpg",
	}

	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Apple is green. I love apple.",
		Price: 120,
		ImgURL: "https://www.paperandtea.com/cdn/shop/articles/Apfel_7ebe153a-a4ac-473a-9217-658894dfc968.jpg?v=1765535477&width=1500",
	}

	prd3 := Product{
		ID: 3,
		Title: "Banna",
		Description: "Banna is Lovely.. I love Banna.",
		Price: 80,
		ImgURL: "https://blog-images-1.pharmeasy.in/blog/production/wp-content/uploads/2021/01/30152155/shutterstock_518328943-1.jpg",
	}

	prd4 := Product{
		ID: 4,
		Title: "Mango",
		Description: "Mango is Lovely.. I love Mango.",
		Price: 160,
		ImgURL: "https://www.paperandtea.com/cdn/shop/articles/Mango_6fb74c95-c9b0-4559-88e8-f542e6d6b18d.jpg?v=1769533193&width=1024",
	}

	prd5 := Product{
		ID: 5,
		Title: "Grape",
		Description: "Grape is purple. I love grape.",
		Price: 200,
		ImgURL: "https://img.imageboss.me/fourwinds/width/425/dpr:2/shop/products/flamegrapes2.jpg?v=1729716550",
	}

	prd6 := Product{
		ID: 6,
		Title: "Pomegranate",
		Description: "Pomegranate is red. I love Pomegranate.",
		Price: 200,
		ImgURL: "https://www.gettystewart.com/wp-content/uploads/2022/11/pomegranate-and-seeds-s.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)


}