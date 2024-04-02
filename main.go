package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HomepageData struct {
	PageTitle string
}

func main() {
	tmpl, err := template.ParseFiles("./views/layout.html")

	if err != nil {
		log.Fatal(err)
	}

	// Templates https://gowebexamples.com/templates/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := HomepageData{
			PageTitle: "VizTreks",
		}

		tmpl.Execute(w, data)
	})

	// Static folder https://gowebexamples.com/static-files/
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist/"))))

	fmt.Println("Serving at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
