package middleware

import "net/http"

// Middleware নামে একটি custom type তৈরি করা হয়েছে
// এটি এমন একটি function type যা একটি http.Handler নেয়
// এবং আরেকটি http.Handler return করে
type Middleware func(http.Handler) http.Handler

// Manager struct তৈরি করা হয়েছে
// এখানে globalMiddlewares নামে একটি slice রাখা হয়েছে
// যেখানে সব global middleware সংরক্ষণ করা হবে
type Manager struct {
	globalMiddlewares []Middleware
}

// NewManager একটি নতুন Manager তৈরি করে return করে
func NewManager() *Manager {
	return &Manager{
		// খালি Middleware slice তৈরি করা হচ্ছে
		globalMiddlewares: make([]Middleware, 0),
	}
}








// Use method এর মাধ্যমে এক বা একাধিক middleware register করা যায়
func (mngr *Manager) Use(middlewares ...Middleware) {

	// নতুন middleware গুলো globalMiddlewares slice এ যোগ করা হচ্ছে
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}












// With method একটি handler এর সাথে middleware যুক্ত করে
// প্রত্যেকটা রিকোয়েস্ট এ কন্ট্রোলাব বা হ্যান্ডেলারে সকল মিডেলওয়্যার কি নিয়ে রিকোয়েস্ট দেয়।
func (mngr *Manager) With(next http.Handler, middlwares ...Middleware) http.Handler {

	// শুরুতে original handler কে n এর মধ্যে রাখা হচ্ছে
	n := next

	// route specific middleware গুলো apply করা হচ্ছে
	for _, middlware := range middlwares {

		// middleware handler কে wrap করছে
		// ফলে নতুন handler আবার n এ সংরক্ষণ হচ্ছে
		n = middlware(n)
	}

	// global middleware গুলো apply করা হচ্ছে
	// for _, globalMiddlware := range mngr.globalMiddlewares {

	// 	// global middleware দিয়েও handler wrap করা হচ্ছে
	// 	n = globalMiddlware(n)
	// }

	// সব middleware apply হওয়ার পর final handler return করা হচ্ছে
	return n
}









// With method একটি handler এর সাথে middleware যুক্ত করে
func (mngr *Manager) WrapMux(next http.Handler, middlwares ...Middleware) http.Handler {

	// শুরুতে original handler কে n এর মধ্যে রাখা হচ্ছে
	n := next

	// route specific middleware গুলো apply করা হচ্ছে
	for _, middlware := range mngr.globalMiddlewares {

		// middleware handler কে wrap করছে
		// ফলে নতুন handler আবার n এ সংরক্ষণ হচ্ছে
		n = middlware(n)
	}

	// সব middleware apply হওয়ার পর final handler return করা হচ্ছে
	return n
}