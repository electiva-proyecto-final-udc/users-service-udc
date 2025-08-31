package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Cami")
}

func main(){
	http.HandleFunc("/", handler)
	fmt.Println("Server listening in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
