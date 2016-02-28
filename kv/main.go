package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	storageFileName := os.TempDir() + "/kv"
	file, err := os.Open(storageFileName)

	k := NewStore(Data{})

	if err == nil {
		k.Read(file)
	}
	file.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			key := r.URL.Path[1:]
			fmt.Fprint(w, k.Filter([]string{key}))
		} else if r.Method == "POST" {
			key := r.URL.Path[1:]
			value, _ := ioutil.ReadAll(r.Body)
			k.Merge(Data{key: string(value)})
			file, _ := os.Create(storageFileName)
			defer file.Close()
			k.Write(file)
		} else {
			http.Error(w, "Method not supported", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
