package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/RomanosTrechlis/archy/graph"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templ, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		type Res struct {
			Title string
		}
		r := Res{"This is a test"}
		templ.Execute(writer, r)
	})

	elements := graph.Get()
	defer graph.Close()
	http.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {

		b, err := json.Marshal(elements)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(b)
	})

	fmt.Println("Server started: localhost:8080")
	http.ListenAndServe(":8080", nil)
}
