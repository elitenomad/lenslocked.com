package main

import "fmt"
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> MY Second First WebPage in GO")
}

func main()  {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on 3000...")
	http.ListenAndServe(":3000", nil)
}
