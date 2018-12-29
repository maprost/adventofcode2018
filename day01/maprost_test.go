package day01

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	input01 = "input01_490.txt"
	input02 = "input02_70357.txt"
)

func TestTask01(t *testing.T) {
	numbers, result := golib.Reads(input01)
	sum := 0

	for _, number := range numbers {
		n, err := strconv.Atoi(strings.TrimLeft(number, "+-"))
		if err != nil {
			fmt.Println(err, " - ", number)
			continue
		}

		if strings.HasPrefix(number, "+") {
			sum += n
		} else {
			sum -= n
		}
	}

	golib.Equal(t, "Sum: ", sum, result)
}

func TestTask02(t *testing.T) {
	numbers, result := golib.Reads(input02)

	set := make(map[int]struct{})
	sum := 0
	more := true
	set[0] = struct{}{}

	for more {
		for _, number := range numbers {
			n, err := strconv.Atoi(strings.TrimLeft(number, "+-"))
			if err != nil {
				fmt.Println(err, " - ", number)
				continue
			}

			if strings.HasPrefix(number, "+") {
				sum += n
			} else {
				sum -= n
			}

			fmt.Println(sum)

			if _, ok := set[sum]; ok {
				golib.Equal(t, " first reaches twice: ", sum, result)
				more = false
				break
			}

			set[sum] = struct{}{}
		}
	}
}
