package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello, world"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    http.Handle("/", r)

    http.ListenAndServe(":12345", nil)
}
