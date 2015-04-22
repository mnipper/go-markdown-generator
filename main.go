package main

import (
  "net/http"
  "fmt"
  "github.com/russross/blackfriday"
)

func main () {
  http.HandleFunc("/markdown", GenerateMarkdown)
  http.Handle("/", http.FileServer(http.Dir("public")))
  fmt.Println(http.ListenAndServe(":3000", nil))
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
  markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
  rw.Write(markdown)
}
