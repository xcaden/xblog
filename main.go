package main

import (
    "net/http"
    "github.com/xcaden/xblog/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler)

    r.PathPrefix("/static/").Handler(
        http.StripPrefix("/static/", http.FileServer(http.Dir("/home/xcaden/Projects/xcaden/xblog/static/"))))

    http.Handle("/", r)
    http.ListenAndServe(":12345", nil)
}
