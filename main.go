package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n" + time.Now().String()))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	log.Println("starting server on port 9001")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":9001", r))
}
