package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/A-D-H-D/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("Error parsing files: %v", err)
		http.Error(w, "There was an error parsing files", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// call the function and pass the filepath
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	// call the execute temp with appropriate filepath
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
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
