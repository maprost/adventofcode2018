package main

import (
	"fmt"

	"github.com/maprost/adventofcode2018/golib"
)

func main() {
	ids := golib.Read("day02/task02/input_oeylbtcxjqnzhgyylfapviusr.txt")

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
			}
		}
	}

}
