package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá Full Cyle Developers!")
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}