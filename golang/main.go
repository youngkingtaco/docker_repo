package main

import ( 
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Hugo's webserver")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServer("0.0.0.0:8080", nil))
}
