package database

var ProductList []Product // ProductList হলো একটি slice যা Product struct এর instance গুলো রাখবে। 
// এটি আমাদের প্রোডাক্ট এর ডাটাবেস হিসেবে কাজ করবে, যেখানে আমরা নতুন প্রোডাক্ট যোগ করতে পারব এবং প্রোডাক্ট এর তথ্য দেখতে পারব।
// P বড় হাতের কারন এটিকে গ্লোভালি সবাই এক্সেস করতে হবে। 


type Product struct {
	ID          int     `json:"id"`          // struct এর বিতরে প্রোপার্টির প্রথম অক্ষর বড় হাতের দিতে হবে। না হলে এই প্রোপার্টি json এ কনভার্ট হবে না।
	Title       string  `json:"title"`       // struct এর প্রথম অক্ষর ছোট হাতের দিলে সেটি প্রাইভেট হয়ে যায়, এটি শুধু মাত্র মেইন প্যাকেজের মধ্যে ব্যবহার করা যাবে, অন্য প্যাকেজে বাবহার করা যাবে না।
	Description string  `json:"description"` // struct এর প্রোপার্টির নাম json এ কনভার্ট করার সময় description হবে।
	Price       float64 `json:"price"`
	// ImgURL      string  `json:"ImgURL_er_poriborthe_ei_nam_dekhabe"` //
	ImgURL string `json:"imageUrl"` //

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


	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)

}
