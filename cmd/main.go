package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Home"))
}

func ProductsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Products"))
}

func ArticlesHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Articles"))
}
