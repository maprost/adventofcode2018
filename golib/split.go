package golib

import (
	"strconv"
	"strings"
)

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
