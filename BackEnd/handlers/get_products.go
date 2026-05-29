package handlers

import (
	"BackEnd/database"
	"BackEnd/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200) // নতুন প্রোডাক্ট এর ডাটা রেসপন্স হিসেবে পাঠানো হচ্ছে।
}
