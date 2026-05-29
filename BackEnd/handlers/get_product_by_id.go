package handlers

import (
	"BackEnd/database"
	"BackEnd/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId") // URL থেকে productId বের করা হচ্ছে, যা প্রোডাক্ট এর id হিসেবে কাজ করবে।

	pId, err := strconv.Atoi(productID) // productId কে integer এ কনভার্ট করা হচ্ছে, কারণ প্রোডাক্ট এর id integer টাইপের হবে।

	if err != nil {
		http.Error(w, "Please give me valid product id", 400)
		return
	}

	for _, product := range database.ProductList { // ekhane _ মানি হচ্ছে যে আমরা index ব্যবহার করব না, শুধুমাত্র product এর ডাটা ব্যবহার করব।

		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}

	}

	util.SendData(w, "Data pai nai", 404)

}
