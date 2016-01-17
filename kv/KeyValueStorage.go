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

func (kv keyValueStorage) filter(keys []string) keyValueStorage {
	res := keyValueStorage{}
	for _, key := range keys {
		if value, ok := kv[key]; ok {
			res[key] = value
		}
	}
	return res
}

func (kv keyValueStorage) String() (res string) {
	for key, value := range kv {
		res += key + " = " + value + "\n"
	}
	return res
}

func (kv keyValueStorage) merge(newkv keyValueStorage) {
	for key, value := range newkv {
		kv[key] = value
	}
}
