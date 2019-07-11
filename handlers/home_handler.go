package handlers

import (
    "os"
    "log"
    "html/template"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    CurrentDir,_ := os.Getwd()
    HtmlDir := CurrentDir + "/static/template/home.html"

    t, err := template.ParseFiles(HtmlDir)
    if err != nil {
        log.Fatal("template.ParseFiles : ", err)
        return
    }

    t.Execute(w, nil)
}
