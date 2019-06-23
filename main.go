package main

import (
    "net/http"
    "github.com/xcaden/xblog/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler)
    http.Handle("/", r)

    http.ListenAndServe(":12345", nil)
}
