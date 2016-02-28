package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, r.URL.Path)
		} else if r.Method == "POST" {
			content, _ := ioutil.ReadAll(r.Body)
			ioutil.WriteFile(r.URL.Path, content, 0644)
		} else {
			http.Error(w, "Method not supported", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
