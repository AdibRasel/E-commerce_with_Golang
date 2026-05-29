package handlers

import (
	"BackEnd/database"
	"BackEnd/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid JSON", 400) // w, err.Error(), 400

		return
	}

	newProduct.ID = len(database.ProductList) + 1 // new product এর id set করা হচ্ছে, যাতে করে প্রতিটি প্রোডাক্ট এর আলাদা আলাদা id থাকে।
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201) // নতুন প্রোডাক্ট এর ডাটা রেসপন্স হিসেবে পাঠানো হচ্ছে।
}
