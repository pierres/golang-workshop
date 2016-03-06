package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_HandlerRoot(t *testing.T) {
	storageFile, _ := ioutil.TempFile(os.TempDir(), "kvtest")
	defer os.Remove(storageFile.Name())
	handler := NewHandler(storageFile.Name())

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Return code was %v", w.Code)
	}
}

func Test_HandlerGetNonExistingKey(t *testing.T) {
	storageFile, _ := ioutil.TempFile(os.TempDir(), "kvtest")
	defer os.Remove(storageFile.Name())
	handler := NewHandler(storageFile.Name())

	req, _ := http.NewRequest("GET", "/some-key", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Body.String() != "" {
		t.Errorf("Return body was %v", w.Body.String())
	}
}

func Test_HandlerStoresValue(t *testing.T) {
	storageFile, _ := ioutil.TempFile(os.TempDir(), "kvtest")
	defer os.Remove(storageFile.Name())
	handler := NewHandler(storageFile.Name())

	key := "some-key"
	content := "some content"

	value := strings.NewReader(content)

	req, _ := http.NewRequest("POST", "/"+key, value)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/"+key, nil)
	handler.ServeHTTP(w, req)
	if w.Body.String() != key+" = "+content+"\n" {
		t.Errorf("Return body was %v", w.Body.String())
	}
}
