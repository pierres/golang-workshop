package main

import (
	"fmt"
	"os"
)

func main() {
	args := Input(os.Args[1:])

	storageFileName := os.TempDir() + "/kv"
	file, err := os.Open(storageFileName)

	k := NewStore(Data{})

	if err == nil {
		k.Read(file)
	}
	file.Close()

	if args.IsWriteStatement() {
		k.Merge(args.Map())

		file, _ := os.Create(storageFileName)
		defer file.Close()

		k.Write(file)
	} else {
		requestedKeys := args.Request()

		fmt.Print(k.Filter(requestedKeys))
	}
}
