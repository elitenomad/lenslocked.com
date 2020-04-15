package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "log"

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> MY Second First WebPage in GO</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>You've requested the contact with name: %s on page \n</h1>", vars["name"])
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> This is first second FAQ page in GO </h1>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "<h1>Not found</h1>")
}

func main()  {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/contact/{name}", contact)
	r.HandleFunc("/faq", faq)

	fmt.Println("Starting server on 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
