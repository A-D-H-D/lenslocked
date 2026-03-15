package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/A-D-H-D/lenslocked/controllers"
	"github.com/A-D-H-D/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

//	func homeHandler(w http.ResponseWriter, r *http.Request) {
//		// call the function and pass the filepath
//		tplPath := filepath.Join("templates", "home.gohtml")
//		executeTemplate(w, tplPath)
//	}

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

	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})

	fmt.Println("Starting web server on port 3000...")
	http.ListenAndServe(":3000", r)
}
