package main

import (
	"fmt"
	"os"
)

func main() {
	args := keyValueInput(os.Args[1:])

	storageFileName := os.TempDir() + "/kv"
	file, err := os.Open(storageFileName)

	k := keyValueStorage{}

	if err == nil {
		k.read(file)
	}
	file.Close()

	if args.isWriteStatement() {
		k.merge(args.getMap())

		file, _ := os.Create(storageFileName)
		defer file.Close()

		k.write(file)
	} else {
		requestedKeys := args.getRequest()

		fmt.Print(k.filter(requestedKeys))
	}
}
