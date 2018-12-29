package main

import (
	"fmt"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
	"github.com/maprost/testbox/must"
)

const (
	input01 = "input01_6916.txt"
	input02 = "input02_oeylbtcxjqnzhgyylfapviusr.txt"
)

func TestTask01(t *testing.T) {
	ids, result := golib.Reads(input01)
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

	checksum := twoTimes * threeTimes
	golib.Equal(t, "CheckSum: ", checksum, result)
}

func TestTask02(t *testing.T) {
	ids, result := golib.Reads(input02)

	for index1, id1 := range ids {
		for index2 := index1; index2 < len(ids); index2++ {
			id2 := ids[index2]
			diffChars := 0
			id := ""

			for i := 0; i < len(id1) && i < len(id2) && diffChars < 2; i++ {
				if id1[i] != id2[i] {
					diffChars++
				} else {
					id += string(id1[i])
				}
			}

			if diffChars == 1 {
				fmt.Println(id1, id2, id)
				must.BeEqual(t, id, result)
			}
		}
	}
}
