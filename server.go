package main

import (
  "net/http"
  "text/template"
)

type Page struct {
  Title string
  Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  count := 2
  page := Page{"Hello World.", count}
  tmpl, err := template.ParseFiles("src/template.html") // ParseFilesを使う
  if err != nil {
    panic(err)
  }
  err = tmpl.Execute(w, page)
  if err != nil {
    panic(err)
  }
}

func main() {
  http.HandleFunc("/", viewHandler) // hello
  http.ListenAndServe(":8080", nil)
}