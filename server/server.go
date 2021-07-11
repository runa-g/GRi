package server

import (
	"fmt"
	"net/http"
	"html/template"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "layout", "root")
}

func ServerLaunch() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
