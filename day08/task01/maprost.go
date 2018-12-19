package main

import (
	"fmt"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

var nodeNumber = 0

type node struct {
	number   int
	children []node
	metadata []int
}

func (n node) String() string {
	return fmt.Sprintf("Node [%d] - c:%d m:%v", n.number, len(n.children), n.metadata)
}

func main() {
	numbers := strings.Split(golib.Read("day08/task01/input_46829.txt")[0], " ")

	root, _ := buildNode(numbers, 0)
	fmt.Println("Sum:", trace(root))
}

func buildNode(numbers []string, index int) (node, int) {
	numberOfChildren := golib.ToInt(numbers[index])
	index++
	numberOfMetadata := golib.ToInt(numbers[index])
	index++

	n := node{
		number:   nodeNumber,
		children: make([]node, 0, numberOfChildren),
		metadata: make([]int, 0, numberOfMetadata),
	}
	nodeNumber++

	for i := 0; i < numberOfChildren; i++ {
		var c node
		c, index = buildNode(numbers, index)
		n.children = append(n.children, c)
	}

	for i := 0; i < numberOfMetadata; i++ {
		n.metadata = append(n.metadata, golib.ToInt(numbers[index]))
		index++
	}
	return n, index
}

func trace(n node) int {
	result := 0
	for i := 0; i < len(n.children); i++ {
		result += trace(n.children[i])
	}
	for i := 0; i < len(n.metadata); i++ {
		result += n.metadata[i]
	}
	return result
}
