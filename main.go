package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
)

type Element struct {
	Data       Data     `json:"data"`
	Position   Position `json:"position"`
	Group      string   `json:"group"`
	Removed    bool     `json:"removed"`
	Selected   bool     `json:"selected"`
	Selectable bool     `json:"selectable"`
	Locked     bool     `json:"locked"`
	Grabbed    bool     `json:"grabbed"`
	Grabbable  bool     `json:"grabbable"`
	Classes    string   `json:"classes"`
}

type Data struct {
	// Node
	Id    string  `json:"id"`
	IdInt int     `json:"idInt"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
	Query bool    `json:"query"`
	Gene  bool    `json:"gene"`

	// Edge
	Source         string  `json:"source"`
	Target         string  `json:"target"`
	Weight         float64 `json:"weight"`
	Group          string  `json:"group"`
	NetworkId      int     `json:"networkId"`
	NetworkGroupId int     `json:"networkGroupId"`
	Intn           bool    `json:"intn"`
	RIntnId        int     `json:"rIntnId"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

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

	b, err := ioutil.ReadFile(filepath.Join("data", "data.json"))

	n1 := Element{
		Data: Data{
			Id:    "612326",
			IdInt: 612326,
			Name:  "DNA1",
			Score: 0.006769776522008331,
			Query: false,
			Gene:  true,
		},
		Position: Position{
			X: 531.9740635094307,
			Y: 464.8210898234145,
		},
		Group:      "nodes",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "fn10273 fn6944 fn9471 fn6284 fn6956 fn6935 fn8147 fn6939 fn6936 fn6949 fn6629 fn7952 fn6680 fn6957 fn8786 fn6676 fn10713 fn7495 fn7500 fn9361 fn6279 fn6278 fn8569 fn7641 fn8568",
	}

	n2 := Element{
		Data: Data{
			Id:    "611408",
			IdInt: 611408,
			Name:  "FEN1",
			Score: 0.006769776522008331,
			Query: false,
			Gene:  true,
		},
		Position: Position{
			X: 531.9740635094307,
			Y: 464.8210898234145,
		},
		Group:      "nodes",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "fn10273 fn6944 fn9471 fn6284 fn6956 fn6935 fn8147 fn6939 fn6936 fn6949 fn6629 fn7952 fn6680 fn6957 fn8786 fn6676 fn10713 fn7495 fn7500 fn9361 fn6279 fn6278 fn8569 fn7641 fn8568",
	}

	n3 := Element{
		Data: Data{
			Id:    "608473",
			IdInt: 608473,
			Name:  "RAD9B",
			Score: 0.006769776522008331,
			Query: false,
			Gene:  true,
		},
		Position: Position{
			X: 751.9740635094307,
			Y: 604.8210898234145,
		},
		Group:      "nodes",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "fn10273 fn6944 fn9471 fn6284 fn6956 fn6935 fn8147 fn6939 fn6936 fn6949 fn6629 fn7952 fn6680 fn6957 fn8786 fn6676 fn10713 fn7495 fn7500 fn9361 fn6279 fn6278 fn8569 fn7641 fn8568",
	}

	e1 := Element{
		Data: Data{
			Source:         "612326",
			Target:         "611408",
			Weight:         0.0055478187,
			Group:          "coexp",
			NetworkId:      1133,
			NetworkGroupId: 18,
			Intn:           true,
			RIntnId:        2,
			Id:             "03",
		},
		Position:   Position{},
		Group:      "edges",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "",
	}

	e2 := Element{
		Data: Data{
			Source:         "612326",
			Target:         "608473",
			Weight:         0.0055478187,
			Group:          "coexp",
			NetworkId:      1133,
			NetworkGroupId: 18,
			Intn:           true,
			RIntnId:        2,
			Id:             "02",
		},
		Position:   Position{},
		Group:      "edges",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "",
	}

	e3 := Element{
		Data: Data{
			Source:         "608473",
			Target:         "611408",
			Weight:         0.0055478187,
			Group:          "coexp",
			NetworkId:      1133,
			NetworkGroupId: 18,
			Intn:           true,
			RIntnId:        2,
			Id:             "01",
		},
		Position:   Position{},
		Group:      "edges",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "",
	}

	elements := make([]Element, 0)
	elements = append(elements, n1)
	elements = append(elements, n2)
	elements = append(elements, n3)
	elements = append(elements, e1)
	elements = append(elements, e2)
	elements = append(elements, e3)

	http.HandleFunc("/data", func(writer http.ResponseWriter, request *http.Request) {
		b, err = json.Marshal(elements)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(b)
	})

	http.ListenAndServe(":8080", nil)
}
