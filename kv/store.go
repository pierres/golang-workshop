package main

import (
	"encoding/json"
	"io"
	"sync"
)

type Data map[string]string

type Store struct {
	content Data
	mutex   *sync.RWMutex
}

func NewStore(data Data) Store {
	kv := Store{}
	kv.content = data
	kv.mutex = &sync.RWMutex{}
	return kv
}

func (kv Store) Read(storage io.Reader) error {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()
	err := json.NewDecoder(storage).Decode(&kv.content)
	return err
}

func (kv Store) Write(storage io.Writer) error {
	data, err := json.Marshal(kv.content)
	if err != nil {
		return err
	}
	kv.mutex.Lock()
	defer kv.mutex.Unlock()
	_, err = storage.Write(data)
	return err
}

func (kv Store) Filter(keys []string) Store {
	if len(keys) == 0 {
		return kv
	}

	res := NewStore(Data{})
	for _, key := range keys {
		if value, ok := kv.content[key]; ok {
			res.content[key] = value
		}
	}
	return res
}

func (kv Store) String() (res string) {
	for key, value := range kv.content {
		res += key + " = " + value + "\n"
	}
	return res
}

func (kv Store) Merge(newkv Data) {
	// TODO: fix unstable iterator
	for key, value := range newkv {
		kv.content[key] = value
	}
}
