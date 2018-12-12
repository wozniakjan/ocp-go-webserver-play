package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("working with hooks!\n"))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
