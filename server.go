// package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./chessfrontend/dist"))
	http.Handle("/", fs)


	fmt.Printf("Starting server on port 8080\n")
	log.Panic(
		http.ListenAndServe(":80", nil),
	)
}