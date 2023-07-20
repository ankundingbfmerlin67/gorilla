package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", YourHandler)
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}
