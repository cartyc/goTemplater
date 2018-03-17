package gotemplater

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var Templates map[string]*template.Template

// Templates Render templates
func LoadTemplates(dir string) {

	Templates = make(map[string]*template.Template)
	base := "templates/layouts/base.html"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() != ".DS_Store" {
			// paths = append(paths, path)
			Templates[info.Name()] = template.Must(template.ParseFiles(base, path))

		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	return
}

// LoadStatic Load static dir
func LoadStatic(dir string) {

	http.Handle("/static/", http.FileServer(http.Dir(dir)))

}
