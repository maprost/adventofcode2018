package day10

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
	"github.com/maprost/testbox/should"
)

const (
	input01 = "input01_LXJFKAXA.txt"
	input02 = "input02_10312.txt"
)

func TestTask01(t *testing.T) {
	golib.ShowDebug = false
	coords, result := golib.Reads(input01)
	run(coords)

	fmt.Println("You should see: ", result)
}

func TestTask02(t *testing.T) {
	coords, result := golib.Reads(input02)
	sec := run(coords)

	golib.Equal(t, "Seconds: ", sec, result)
}

// ================== tests =============================

func TestConvertPoints(t *testing.T) {
	should.BeEqual(t, convertPoint("position=<-51359, -51442> velocity=< 5,  6>"), &point{
		x:  -51359,
		y:  -51442,
		mx: 5,
		my: 6,
	})
}

// ================== program =============================

type point struct {
	x  int
	y  int
	mx int
	my int
}

func (p *point) shift() {
	p.x += p.mx
	p.y += p.my
}

func (p *point) revertShift() {
	p.x -= p.mx
	p.y -= p.my
}

type points struct {
	points []*point
	maxX   int
	minX   int
	maxY   int
	minY   int
}

func (ps *points) add(p *point) {
	ps.points = append(ps.points, p)
	ps.calcRanges(p)
}

func (ps *points) calcRanges(p *point) {
	if ps.minX > p.x {
		ps.minX = p.x
	}
	if ps.maxX < p.x {
		ps.maxX = p.x
	}
	if ps.minY > p.y {
		ps.minY = p.y
	}
	if ps.maxY < p.y {
		ps.maxY = p.y
	}
}

func (ps *points) shift() {
	ps.resetBox()

	for _, p := range ps.points {
		p.shift()
		ps.calcRanges(p)
	}
}

func (ps *points) revertShift() {
	ps.resetBox()

	for _, p := range ps.points {
		p.revertShift()
		ps.calcRanges(p)
	}
}

func (ps *points) resetBox() {
	ps.maxY = 0
	ps.minY = 0
	ps.maxX = 0
	ps.minX = math.MaxInt32
}

func (ps *points) box() int {
	return (golib.IntAbs(ps.minX) + golib.IntAbs(ps.maxX)) * (golib.IntAbs(ps.minY) + golib.IntAbs(ps.maxY))
}

func (ps *points) show() bool {
	xOffset := ps.minX * (-1)
	yOffset := ps.minY * (-1)

	xSize := xOffset + ps.maxX + 1
	ySize := yOffset + ps.maxY + 1

	coords := make([][]bool, ySize)
	for i := range coords {
		coords[i] = make([]bool, xSize)
	}

	golib.Debugln("Setup: x:", xSize, " y:", ySize)

	for _, p := range ps.points {
		x := p.x + xOffset
		y := p.y + yOffset
		coords[y][x] = true
	}

	golib.Debugln("build output")

	for y := range coords {
		for x := range coords[y] {
			if coords[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return true
}

func run(coords []string) int {

	sky := points{points: make([]*point, 0, len(coords))}
	sky.resetBox()

	for _, coord := range coords {
		sky.add(convertPoint(coord))
	}

	oldBox := sky.box()
	sky.shift()
	i := 0

	for oldBox > sky.box() {
		oldBox = sky.box()
		sky.shift()
		i++
	}

	sky.revertShift()
	sky.show()
	return i
}

// position=< y,  x> velocity=< my,  mx>
func convertPoint(in string) *point {
	in = strings.Replace(in, "position=<", "", 1)
	in = strings.Replace(in, ",", ";", -1)
	in = strings.Replace(in, "> velocity=<", ";", 1)
	in = strings.Replace(in, ">", "", 1)
	in = strings.Replace(in, " ", "", -1)

	inArr := strings.Split(in, ";")
	return &point{
		x:  golib.ToInt(inArr[0]),
		y:  golib.ToInt(inArr[1]),
		mx: golib.ToInt(inArr[2]),
		my: golib.ToInt(inArr[3]),
	}
}
