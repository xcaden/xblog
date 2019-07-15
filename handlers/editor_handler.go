package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func EditorHandler(w http.ResponseWriter, r *http.Request) {
	CurrentDir, _ := os.Getwd()
	HtmlDir := CurrentDir + "/static/template/editor.html"

	t, err := template.ParseFiles(HtmlDir)
	if err != nil {
		log.Fatal("template.ParseFiles : ", err)
		return
	}

	t.Execute(w, nil)
}
