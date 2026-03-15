package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

/* Must functions in go */
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFs(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("Parse Fs template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("Parsing template: %w", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-type", "text/html")
	err := t.htmlTpl.Execute(w, nil)

	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
