package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Something bad happened!"))
  })

  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":8080", nil)
}
