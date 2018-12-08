package golib

import (
	"io/ioutil"
	"strings"
)

func Read(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
