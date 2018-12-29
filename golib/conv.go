package golib

import (
	"math"
	"strconv"
)

func ToInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func IntAbs(i int) int {
	return int(math.Abs(float64(i)))
}
