package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/A-D-H-D/lenslocked/controllers"
	"github.com/A-D-H-D/lenslocked/templates"
	"github.com/A-D-H-D/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// r.Get("/", controllers.StaticHandler(
	// 	views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	/** changed from previous one to the one below */

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFs(templates.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFs(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFs(templates.FS, "faq.gohtml"))))

	r.Get("/shame", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "shame.gohtml")))))

	// r.NotFound(func(w http.ResponseWriter, r *http.Request) {
	// 	http.Error(w, "Not found", http.StatusNotFound)
	// })

	r.NotFound(controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "404.gohtml")))))

	fmt.Println("Starting web server on port 3000...")
	http.ListenAndServe(":3000", r)
}
