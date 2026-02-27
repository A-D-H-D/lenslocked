package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Good Server	</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch email me at <a href=\"mailto:ulqiora@gmail.com\">ulqiora@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	fmt.Fprint(w, `<h1>FAQ  page</h1>
	<ul>
		<li>Is there a free version? <b>You bet hombre :) </li>
		<li>How do you contact support? <b>:)<a href=\"mailto:ulqiora@gmail.com\">ulqiora@gmail.com</a> </li>

	</ul>

	`)

}

// func pathHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)

// 	default:
// 		// to do handle page not found errer
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})

	fmt.Println("Starting web server on port 3000...")
	http.ListenAndServe(":3000", r)
}
