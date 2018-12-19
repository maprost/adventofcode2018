package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

type ringNode struct {
	value int
	left  *ringNode
	right *ringNode
}

func main() {
	in := golib.Read("day09/task01/input_439089.txt")[0]
	inArr := strings.Split(in, " players; last marble is worth ")

	player := golib.ToInt(inArr[0])
	max := golib.ToInt(strings.Split(inArr[1], " points")[0])

	fmt.Printf("Setting: Player(%d) Max(%d)\n", player, max)

	playerScore := make(map[int]int)

	root := &ringNode{value: 0}
	root.left = root
	root.right = root

	currentMarble := root
	currentPlayer := 1
	for m := 1; m <= max; m++ {
		if m%23 == 0 {
			// delete ringNode (-7) and get score of the node

			currentMarble = currentMarble.left.left.left.left.left.left.left
			score := m + currentMarble.value

			currentMarble.left.right = currentMarble.right
			currentMarble.right.left = currentMarble.left

			currentMarble = currentMarble.right

			pScore := playerScore[currentPlayer]
			playerScore[currentPlayer] = score + pScore
		} else {
			// add ringNode (+1)

			currentMarble = currentMarble.right
			newMarble := &ringNode{
				value: m,
				left:  currentMarble,
				right: currentMarble.right,
			}
			currentMarble.right.left = newMarble
			currentMarble.right = newMarble

			currentMarble = newMarble
		}

		//fmt.Printf("[%d] %s c[%d]\n", currentPlayer, ring(root), currentMarble.value)

		currentPlayer++
		if currentPlayer > player {
			currentPlayer = 1
		}
	}

	maxScore := 0
	for _, score := range playerScore {
		if maxScore < score {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}

func ring(root *ringNode) string {
	str := strconv.Itoa(root.value) + " "

	currentMarble := root.right
	for currentMarble.value != root.value {
		str += strconv.Itoa(currentMarble.value) + " "
		currentMarble = currentMarble.right
	}

	return str
}
