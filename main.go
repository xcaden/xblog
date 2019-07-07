package main

import (
    "os"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/xcaden/xblog/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler)
    r.HandleFunc("/ajax/{name}", handlers.AjaxHandler)

    CurrentDir,_ := os.Getwd()
    StaticDir := CurrentDir + "/static/"
    r.PathPrefix("/static/").Handler(
        http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
