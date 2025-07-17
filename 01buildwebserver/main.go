package main

import (
	"fmt"
	"log"
	"net/http"
)

// Func form Handler

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address=%s\n", address)
}

// / FUNC HELLO HANDLEER
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supprted", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hellow")
}

func main() {
	fmt.Println("WELCOME TO MY WEB SERVER")

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server ad port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
