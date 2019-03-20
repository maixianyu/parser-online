package main

import (
	//"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"log"
)

var editorTmpl *template.Template

type ParsePage struct {
	Title string
	Intro string
	Code  string
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/clojure", cljHandler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r

}

func main() {
	var err error
	editorTmpl, err = template.ParseFiles("assets/editor.html")
	if err != nil {
		log.Fatal(err)
	}

	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func cljHandler(w http.ResponseWriter, r *http.Request) {
	clj := ParsePage {
		Title: "clojure parse-online",
		Intro: "clojure is a function lang",
		Code:  "()",
	}
	editorTmpl.Execute(w, clj)
}

