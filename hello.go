package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
func main() {
	fmt.Printf("!oG ,olleH\n")
	fmt.Printf(stringutil.Reverse("!oG ,olleH") + "\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Hello, %q", r.Header)
	})

	fmt.Printf("Listening on 8080\n")
	http.ListenAndServe(":8080", nil)
}
*/

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/products", ProductsHandler)
	router.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", router)

	// Subrouter
	s := router.PathPrefix("/products").Subrouter()
	s.HandleFunc("/products/", ProductsHandler)
	s.HandleFunc("/products/{key}", ProductHandler)

	router.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("article")
	url, err := router.Get("article").URL("category", "technology", "id", "42")

	if err != nil {
		fmt.Printf("Printing error, %q", err)
	} else {
		fmt.Printf("Printing custom url, %q\n", url)
	}
	fmt.Printf("Listening on 8080\n")

	http.ListenAndServe(":8080", nil)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func ProductsHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "products")
}

func ProductHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "product without s")
}

func ArticlesHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "articles")
}

func ArticleHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "article without s")
}
