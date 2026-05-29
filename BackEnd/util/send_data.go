package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)     // এখানে data interface{} দেয়া হয়েছে, যাতে করে যেকোনো ধরনের ডাটা পাঠানো যায়, যেমন struct, string, int ইত্যাদি।
	encoder := json.NewEncoder(w) // statusCode int দেয়া হয়েছে, যাতে করে যেকোনো ধরনের স্ট্যাটাস কোড পাঠানো যায়, যেমন 200, 201, 400 ইত্যাদি।
	encoder.Encode(data)          // encoder দ্বারা data কে json format এ convert করা হচ্ছে এবং w এর মাধ্যমে রেসপন্স হিসেবে পাঠানো হচ্ছে।
}
