package main

import (
	"strings"
)

type keyValueInput []string

func (input keyValueInput) isWriteStatement() bool {
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

func (input keyValueInput) getRequest() []string {
	return input
}

func (input keyValueInput) getMap() map[string]string {
	keyValues := make(map[string]string)
	for _, line := range input {
		splitLine := strings.Split(line, "=")
		keyValues[splitLine[0]] = splitLine[1]
	}
	return keyValues
}
