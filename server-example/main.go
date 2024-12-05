package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error while parsing the form")
		return
	}
	fmt.Fprintf(w, "Form uploaded successfully !!!")
}
func helloHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method id not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello call is done here !!!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer) //to handle the root folder directly returning the index.html
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHanlder)

	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
