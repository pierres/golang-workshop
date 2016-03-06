package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Handler struct {
	storageFileName string
	k               Store
}

func NewHandler(storageFileName string) Handler {
	h := Handler{}

	h.storageFileName = storageFileName
	file, err := os.Open(h.storageFileName)
	h.k = NewStore(Data{})
	if err == nil {
		h.k.Read(file)
	}
	file.Close()

	return h
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		key := r.URL.Path[1:]
		fmt.Fprint(w, h.k.Filter([]string{key}))
	} else if r.Method == "POST" {
		key := r.URL.Path[1:]
		value, _ := ioutil.ReadAll(r.Body)
		h.k.Merge(Data{key: string(value)})
		file, _ := os.Create(h.storageFileName)
		defer file.Close()
		h.k.Write(file)
	} else {
		http.Error(w, "Method not supported", http.StatusInternalServerError)
	}
}

func main() {
	handler := NewHandler(os.TempDir() + "/kv")
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
