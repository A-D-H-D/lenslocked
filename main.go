package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("content-Type", "text/html")

	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Error parsing files: %v", err)
		http.Error(w, "There was an error parsing files", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("There was an error executing templates")
		http.Error(w, "Error executing templates", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// call the function and pass the filepath
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	w.Header().Set("content-Type", "text/html")

	// second method to parsefiles
	tplPath := filepath.Join("templates", "contact.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Error parsing the template: %v", err)
		http.Error(w, "Error parsing the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
=======
	// call the execute temp with appropriate filepath
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
>>>>>>> contact-templete
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
