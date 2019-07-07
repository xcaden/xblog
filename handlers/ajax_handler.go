package handlers

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func AjaxHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    fmt.Printf("%v",vars)
}
