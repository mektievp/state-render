package main

import (
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("/home/phillip/develop/state-render"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
