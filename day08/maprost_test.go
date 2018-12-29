package day08

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	input01 = "input01_46829.txt"
	input02 = "input02_37450.txt"
)

func TestTask01(t *testing.T) {
	input, result := golib.Reads(input01)
	numbers := strings.Split(input[0], " ")

	root, _ := buildNode(numbers, 0)
	golib.Equal(t, "Sum:", calcMetadata(root), result)
}

func TestTask02(t *testing.T) {
	input, result := golib.Reads(input02)
	numbers := strings.Split(input[0], " ")

	root, _ := buildNode(numbers, 0)
	golib.Equal(t, "Sum:", calcMetaDataOfEmptyNodes(root), result)
}

var nodeNumber = 0

type node struct {
	number   int
	children []node
	metadata []int
}

func (n node) String() string {
	return fmt.Sprintf("Node [%d] - c:%d m:%v", n.number, len(n.children), n.metadata)
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

func calcMetadata(n node) int {
	result := 0
	for i := 0; i < len(n.children); i++ {
		result += calcMetadata(n.children[i])
	}
	for i := 0; i < len(n.metadata); i++ {
		result += n.metadata[i]
	}
	return result
}

func calcMetaDataOfEmptyNodes(n node) int {
	fmt.Println(n.String())

	result := 0

	if len(n.children) == 0 {
		for i := 0; i < len(n.metadata); i++ {
			result += n.metadata[i]
		}

	} else {
		for i := 0; i < len(n.metadata); i++ {
			childIndex := n.metadata[i] - 1

			if childIndex >= 0 && childIndex < len(n.children) {
				result += calcMetaDataOfEmptyNodes(n.children[childIndex])
			}
		}
	}

	return result
}
