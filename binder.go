package binder

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Binder struct {
	Router    *mux.Router
	templates *template.Template
	baseDir   string
}

// function new used to start the binder
// baseTempalte can be a file or a pattern like "templates/*html"
func New(baseTemplate string) *Binder {
	baseDir := filepath.Dir(baseTemplate)
	pattern := filepath.Join(baseDir, "*html")

	tmpl := template.Must(template.ParseGlob(pattern))

	return &Binder{
		Router:    mux.NewRouter(),
		templates: tmpl,
		baseDir:   baseDir,
	}
}

// function NewRoute used to create a route based on a HTML template
func (b *Binder) NewRoute(path string, templateName string) {
	b.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		err := b.templates.ExecuteTemplate(w, templateName, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("rendering error %s: %v", templateName, err), http.StatusInternalServerError)
		}
	})
}

// function NewRouteData used to create a route w dynamic data
func (b *Binder) NewRouteData(path string, templateName string, dataFunc func(*http.Request) any) {
	b.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		data := dataFunc(r)
		err := b.templates.ExecuteTemplate(w, templateName, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("rendering error %s: %v", templateName, err), http.StatusInternalServerError)
		}
	})
}

// function serve used to start the server with a chosen port
func (b *Binder) Serve(addr string) error {
	fmt.Printf("Server started on https://localhost%s\n", addr)
	return http.ListenAndServe(addr, b.Router)
}
