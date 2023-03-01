package main

import (
    "fmt""net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to my website!")
        fmt.Fprintln(w, r.URL.Query().Get("name"))
        fmt.Fprintln(w, r.FormValue("email"))})
        fs := http.FileServer(http.Dir("static/"))
        http.Handle("/static/", http.StripPrefix("/static/", fs))
        http.ListenAndServe(":8080", nil)
}