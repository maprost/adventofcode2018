package main

import (
	"fmt"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	a = 'a'
	A = 'A'
)

var (
	offset = uint8(a - A)
)

func main() {
	code := golib.Read("day05/task01/input_9370.txt")[0]

	for {
		replacer := make([]string, 0)
		for i := 0; i < len(code)-1; i++ {
			if code[i] == (code[i+1]+offset) || code[i]+offset == code[i+1] {
				replacer = append(replacer, string(code[i])+string(code[i+1]))
			}
		}

		if len(replacer) == 0 {
			break
		}

		for _, r := range replacer {
			code = strings.Replace(code, r, "", 1)
		}
	}

	fmt.Println("Len(code): ", len(code))
}
