package gotemplater

import (
	"github.com/oxtoacart/bpool"
	"path/filepath"
	"html/template"
	"log"
	"net/http"
	"fmt"
)

/*
Adapted From
https://hackernoon.com/golang-template-2-template-composition-and-how-to-organize-template-files-4cb40bcdf8f6
*/

var templates map[string]*template.Template
var bufpool *bpool.BufferPool


type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var templateConfig TemplateConfig

func LoadConfiguration(main string, subtemplates string) {
	templateConfig.TemplateLayoutPath = subtemplates
	templateConfig.TemplateIncludePath = main
}

func LoadStatic(dir string){

	http.Handle("/img/", http.FileServer(http.Dir(dir)))
	http.Handle("/css/", http.FileServer(http.Dir(dir)))
	http.Handle("/js/", http.FileServer(http.Dir(dir)))

}

func LoadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}


	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	fmt.Println("Template")
	fmt.Println(tmpl)

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}