#Go template loader

go get github.com/cartyc/gotemplater

## functions

LoadConfiguration(main string, subtemplates string)


LoadStatic(dir string)


LoadTemplates()


RenderTemplate(w http.ResponseWriter, name string, data interface{})


Example Usage

```
func test(w http.ResponseWriter, r *http.Request) {
	gotemplater.RenderTemplate(w, "test.html", nil)
}


func main() {
	gotemplater.LoadStatic("static")
	gotemplater.LoadConfiguration("templates/", "templates/layout/")
	gotemplater.LoadTemplates()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login/", test)
	server.ListenAndServe()
}
```