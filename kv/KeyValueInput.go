package main

import (
	"strings"
)

type keyValueInput []string

func (input keyValueInput) isWriteStatement() bool {
	for _, arg := range input {
		if strings.Count(arg, "=") != 1 {
			return false
		}
	}
	return true
}

func (input keyValueInput) isReadStatement() bool {
	return !input.isWriteStatement()
}

func (input keyValueInput) getRequest() []string {
	// TODO: validate
	return input
}

func (input keyValueInput) getMap() map[string]string {
	// TODO: validate
	keyValues := make(map[string]string)
	for _, line := range input {
		splitLine := strings.Split(line, "=")
		keyValues[splitLine[0]] = splitLine[1]
	}
	return keyValues
}
