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
	code := golib.Read("day05/task02/input_6390.txt")[0]

	minSize := len(code)

	for i := 'a'; i <= 'z'; i++ {
		// prepare code
		rCode := strings.Replace(code, string(i), "", -1)
		rCode = strings.Replace(rCode, string(i-int32(offset)), "", -1)

		// reduce
		for {
			replacer := make([]string, 0)
			for i := 0; i < len(rCode)-1; i++ {
				if rCode[i] == (rCode[i+1]+offset) || rCode[i]+offset == rCode[i+1] {
					replacer = append(replacer, string(rCode[i])+string(rCode[i+1]))
				}
			}

			if len(replacer) == 0 {
				break
			}

			for _, r := range replacer {
				rCode = strings.Replace(rCode, r, "", 1)
			}
		}

		// check
		if minSize > len(rCode) {
			minSize = len(rCode)
		}
	}

	fmt.Println("min Len(rCode): ", minSize)
}
