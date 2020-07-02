package api

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Text string
}

type Response struct {
	Text string
}

func Test() {
	tmpl := template.Must(template.ParseFiles("create.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		info := Data{
			Text: r.FormValue("text"),
		}

		fmt.Println(info)

		_ = AddtoDB(info)

	})

	http.ListenAndServe(":8080", nil)
}

// adds the
func AddtoDB(info Data) Response {
	return Response{"localhost:8080/"}
}
