package day05

import (
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	a       = 'a'
	A       = 'A'
	input01 = "input01_9370.txt"
	input02 = "input02_6390.txt"
)

var (
	offset = uint8(a - A)
)

func TestTask01(t *testing.T) {
	codes, result := golib.Reads(input01)
	code := codes[0]

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
	golib.Equal(t, "Len(code): ", len(code), result)
}

func TestTask02(t *testing.T) {
	codes, result := golib.Reads(input02)
	code := codes[0]
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
	golib.Equal(t, "min Len(rCode): ", minSize, result)
}
