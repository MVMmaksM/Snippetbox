package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal server error", 500)
		return
	}

	err = ts.Execute(writer, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal server error", 500)
	}
}

func about(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Страница about"))
}

func showSnippet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, "Выбранная заметка с id %d", id)
}

func createSnippet(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Метод запрещен!", 405)
		return
	}
	writer.Write([]byte(`{"name":"Test"}`))
}
