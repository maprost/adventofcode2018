package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("day01/task01/go/numbers.txt") // just pass the file name
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(string(b), "\n")
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
