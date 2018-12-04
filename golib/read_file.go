package golib

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Read(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}

func SplitFabricInput(in string) (id, leftEdge, topEdge, wide, height int) {
	var err error
	inArr := strings.Split(in, " ")

	id, err = strconv.Atoi(strings.TrimLeft(inArr[0], "#"))
	if err != nil {
		panic(err)
	}

	edgePoints := strings.Split(inArr[2], ",")
	leftEdge, err = strconv.Atoi(edgePoints[0])
	if err != nil {
		panic(err)
	}

	topEdge, err = strconv.Atoi(strings.TrimRight(edgePoints[1], ":"))
	if err != nil {
		panic(err)
	}

	size := strings.Split(inArr[3], "x")
	wide, err = strconv.Atoi(size[0])
	if err != nil {
		panic(err)
	}

	height, err = strconv.Atoi(size[1])
	if err != nil {
		panic(err)
	}

	return
}
