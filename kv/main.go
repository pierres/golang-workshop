package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
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
		h.Get(w, r, httprouter.Params{httprouter.Param{"key", key}})
	} else if r.Method == "POST" {
		key := r.URL.Path[1:]
		h.Post(w, r, httprouter.Params{httprouter.Param{"key", key}})
	} else {
		http.Error(w, "Method not supported", http.StatusInternalServerError)
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	fmt.Fprint(w, h.k.Filter([]string{key}))
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	value, _ := ioutil.ReadAll(r.Body)
	h.k.Merge(Data{key: string(value)})
	file, _ := os.Create(h.storageFileName)
	defer file.Close()
	h.k.Write(file)
}

func main() {
	mux := httprouter.New()
	handler := NewHandler(os.TempDir() + "/kv")
	mux.Handler("GET", "/:key", handler)
	mux.Handler("POST", "/:key", handler)
	panic(http.ListenAndServe(":8080", mux))
}
