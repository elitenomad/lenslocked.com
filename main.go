package main

import "fmt"
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> MY Second First WebPage in GO</h1>")
}

func contactHanderFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> This is first second contact page in GO </h1>")
}

func faqHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> This is first second FAQ page in GO </h1>")
}

func main()  {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contact", contactHanderFunc)
	http.HandleFunc("/faq", faqHandlerFunc)

	fmt.Println("Starting server on 3000...")
	http.ListenAndServe(":3000", nil)
}
