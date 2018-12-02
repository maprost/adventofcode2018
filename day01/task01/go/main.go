package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

func main() {
	numbers := golib.Read("day01/task01/go/numbers.txt")
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

	fmt.Println("Summe: ", sum)
}
