package main

import (
	"strings"
)

type Input []string

func (input Input) IsWriteStatement() bool {
	if len(input) == 0 {
		return false
	}
	for _, arg := range input {
		if strings.Count(arg, "=") != 1 {
			return false
		}
	}
	return true
}

func (input Input) Request() []string {
	return input
}

func (input Input) Map() map[string]string {
	keyValues := make(map[string]string)
	for _, line := range input {
		splitLine := strings.Split(line, "=")
		keyValues[splitLine[0]] = splitLine[1]
	}
	return keyValues
}
