package handlers

import (
	"html/template"
	"net/http"
)

//HomeHandler : Handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/home.html", "tmpl/dash.html")
	t.Execute(w, nil)
}
