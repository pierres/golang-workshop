package main

import (
	"fmt"
	"os"
)

func main() {
	args := keyValueInput(os.Args[1:])

	storageFileName := os.TempDir() + "/kv"

	if args.isWriteStatement() {
		k := keyValueStorage(args.getMap())

		file, _ := os.Create(storageFileName)
		defer file.Close()

		k.write(file)
	} else {
		requestedKeys := args.getRequest()
		file, _ := os.Open(storageFileName)
		defer file.Close()

		k := keyValueStorage{}
		k.read(file)
		for _, key := range requestedKeys {
			fmt.Println(key + "=" + k[key])
		}
	}
}
