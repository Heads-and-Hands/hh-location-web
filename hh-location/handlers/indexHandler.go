package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexViewData struct {
	text string
}

var IndexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	data := IndexViewData{
		text: "Hello",
	}

	wd, err := os.Getwd()
	if err != nil {	log.Fatal(err) }
	tmpl, err := template.ParseFiles(wd + "/templates/index.html")
	if err != nil { log.Fatal(err) }

	tmpl.Execute(w, data)
})
