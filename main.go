package main

import (
	"fmt"
	"log"
	"net/http"
)

func holaFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo!")
}

func main() {
	canapp := http.FileServer(http.Dir("./static"))

	http.Handle("/", canapp)

	log.Println("Running app on 4000 port")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
