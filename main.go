package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl, err := template.ParseFiles("layout.html")

	if err != nil {
		log.Fatal(err)
	}

	// Templates https://gowebexamples.com/templates/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "TODO List",
			Todos: []Todo{
				{Title: "Eat dinner", Done: false},
				{Title: "Drink coffee", Done: true},
				{Title: "Do laundry", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})

	// Static folder https://gowebexamples.com/static-files/
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serving at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
