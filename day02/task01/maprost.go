package main

import (
	"fmt"

	"github.com/maprost/adventofcode2018/golib"
)

func main() {
	ids := golib.Read("day02/task01/input_6916.txt")
	twoTimes := 0
	threeTimes := 0

	for _, id := range ids {
		counter := make(map[int32]int)

		for _, c := range id {
			n, _ := counter[c]
			counter[c] = n + 1
		}

		// check for count 2
		for _, count := range counter {
			if count == 2 {
				twoTimes++
				break
			}
		}

		// check for count 3
		for _, count := range counter {
			if count == 3 {
				threeTimes++
				break
			}
		}
	}

	fmt.Printf("CheckSum: %d*%d = %d", twoTimes, threeTimes, twoTimes*threeTimes)
}
