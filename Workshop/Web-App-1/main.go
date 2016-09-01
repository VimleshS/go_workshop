package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", serverTemplate)
	http.HandleFunc("/save", save)
	go http.ListenAndServe(":5000", nil)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)

	log.Println("Server is started")
}

func save(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	password := r.FormValue("password")
	log.Printf("%s --- %s\n", name, password)
}

func serverTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("template", "layout.html")
	tmpl := template.Must(template.New("layout").ParseFiles(lp))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
