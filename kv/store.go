package main

import (
	"encoding/json"
	"io"
)

type Store map[string]string

func (kv *Store) Read(storage io.Reader) error {
	err := json.NewDecoder(storage).Decode(kv)
	return err
}

func (kv *Store) Write(storage io.Writer) error {
	data, err := json.Marshal(kv)
	if err != nil {
		return err
	}
	_, err = storage.Write(data)
	return err
}

func (kv Store) Filter(keys []string) Store {
	if len(keys) == 0 {
		return kv
	}

	res := Store{}
	for _, key := range keys {
		if value, ok := kv[key]; ok {
			res[key] = value
		}
	}
	return res
}

func (kv Store) String() (res string) {
	for key, value := range kv {
		res += key + " = " + value + "\n"
	}
	return res
}

func (kv Store) Merge(newkv Store) {
	for key, value := range newkv {
		kv[key] = value
	}
}
