package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

func main() {
	numbers := golib.Read("day01/task02/go/numbers.txt")

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
				fmt.Println(" first reaches ", sum, " twice")
				more = false
				break
			}

			set[sum] = struct{}{}
		}
	}

}
