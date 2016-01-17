package main

import (
	"encoding/json"
	"io"
)

type keyValueStorage map[string]string

func (kv *keyValueStorage) read(storage io.Reader) error {
	err := json.NewDecoder(storage).Decode(kv)
	return err
}

func (kv *keyValueStorage) write(storage io.Writer) error {
	data, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	_, err = storage.Write(data)
	return err
}
