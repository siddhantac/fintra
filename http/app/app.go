package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

/*
func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	render(w, r, []byte(newTransaction))
}

func ListAllTransaction(w http.ResponseWriter, r *http.Request) {
	render(w, r, []byte(newTransaction))
}

func render(w http.ResponseWriter, r *http.Request, page []byte) {
	_, err := w.Write(page)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}
*/

var htmlFiles = []string{
	"./templates/index.html",
	"./templates/newTransaction.html",
}

type Page struct {
	Title template.HTML
}

func Index() http.HandlerFunc {
	templates := template.Must(template.ParseFiles(htmlFiles...))
	pg := Page{Title: template.HTML("My test heading")}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "base", &pg)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse template: %v", err), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}

// Public serves static assets such as CSS and JavaScript to clients.
func Public(w http.ResponseWriter, r *http.Request) {
	p, _ := os.Getwd()
	fmt.Println(">>> public", p)
	http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
}
