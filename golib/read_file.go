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

func Reads(file string) ([]string, string) {
	return Read(file), strings.Split(strings.Split(file, "_")[1], ".txt")[0]
}
