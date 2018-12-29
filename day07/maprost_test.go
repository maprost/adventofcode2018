package day07

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
	"github.com/maprost/testbox/must"
)

const (
	input01 = "input01_FHICMRTXYDBOAJNPWQGVZUEKLS.txt"
)

type Node struct {
	name     string
	parents  []*Node
	children []*Node
}

func (n *Node) allParentsDone(nodesInUse map[string]struct{}) bool {
	for _, p := range n.parents {
		if _, ok := nodesInUse[p.name]; !ok {
			return false
		}
	}
	return true
}

type NodeMap map[string]*Node

func (n *NodeMap) get(name string) *Node {
	node, ok := (*n)[name]
	if !ok {
		node = &Node{
			name:     name,
			children: make([]*Node, 0),
			parents:  make([]*Node, 0),
		}
		(*n)[name] = node
	}
	return node
}

type Nodes []*Node

func (n *Nodes) sort(finishNodes map[string]struct{}) {
	sort.Slice(*n, func(i, j int) bool {
		if (*n)[i].allParentsDone(finishNodes) == false {
			return false
		}

		if (*n)[j].allParentsDone(finishNodes) == false {
			return true
		}

		return (*n)[i].name < (*n)[j].name
	})
}

func TestTask01(t *testing.T) {
	orders, result := golib.Reads(input01)
	nodes := make(NodeMap)

	for _, order := range orders {
		target, source := split(order)

		tNode := nodes.get(target)
		sNode := nodes.get(source)

		tNode.children = append(tNode.children, sNode)
		sNode.parents = append(sNode.parents, tNode)
	}

	// find start nodes
	possibleNodes := make(Nodes, 0)
	nodesInUse := make(map[string]struct{})
	finishNodes := make(map[string]struct{})

	for _, node := range nodes {
		if len(node.parents) == 0 {
			possibleNodes = append(possibleNodes, node)
			nodesInUse[node.name] = struct{}{}
			fmt.Println("StartNode: ", node.name)
		}
	}

	// trace
	possibleNodes.sort(finishNodes)
	route := ""

	for len(possibleNodes) > 0 {
		currentNode := possibleNodes[0]
		route += currentNode.name

		possibleNodes = possibleNodes[1:]
		finishNodes[currentNode.name] = struct{}{}

		for _, c := range currentNode.children {
			if _, ok := nodesInUse[c.name]; !ok {
				nodesInUse[c.name] = struct{}{}
				possibleNodes = append(possibleNodes, c)
			}
		}

		possibleNodes.sort(finishNodes)
	}

	fmt.Println("Route: ", route)
	must.BeEqual(t, route, result)
}

func split(in string) (target, source string) {
	in = strings.Replace(in, "Step ", "", 1)
	in = strings.Replace(in, " must be finished before step ", "-", 1)
	in = strings.Replace(in, " can begin.", "", 1)

	inArr := strings.Split(in, "-")
	return inArr[0], inArr[1]
}

func print(nodes []*Node) {
	for _, n := range nodes {
		fmt.Print(n.name, ",")
	}
}
