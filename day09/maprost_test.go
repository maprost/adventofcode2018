package day09

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	input01 = "input01_439089.txt"
	input02 = "input02_3668541094.txt"
)

func TestTask01(t *testing.T) {
	doIt(t, input01)
}

func TestTask02(t *testing.T) {
	doIt(t, input02)
}

// ================== program =============================

type ringNode struct {
	value int
	left  *ringNode
	right *ringNode
}

func doIt(t testing.TB, file string) {
	input, result := golib.Reads(file)
	in := input[0]
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
	golib.Equal(t, "MaxScore: ", maxScore, result)
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
