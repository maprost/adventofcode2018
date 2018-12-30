package day12

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
	"github.com/maprost/testbox/must"
)

const (
	edge    = 150
	input01 = "input01_3410.txt"
	input02 = "input02_4000000001480.txt"
)

func TestTask01(t *testing.T) {
	in, result := golib.Reads(input01)
	input := convertInput(in)

	sum := grow(input, 20)
	golib.Equal(t, "Sum: ", sum, result)
}

func TestTask02(t *testing.T) {
	in, result := golib.Reads(input02)
	input := convertInput(in)

	changeTime := 114
	sum := grow(input, 114)
	must.BeEqual(t, sum, 10600)
	must.BeEqual(t, grow(input, 115), 10680) // +80
	must.BeEqual(t, grow(input, 116), 10760) // +80
	must.BeEqual(t, grow(input, 117), 10840) // +80
	must.BeEqual(t, grow(input, 127), 11640) // +800

	// guessing: 80 * (50000000000-114) + 10600
	golib.Equal(t, "Sum:", int((80*(50000000000-changeTime))+sum), result)
}

// ================== tests =============================

func TestConverter(t *testing.T) {
	input := []string{
		"initial state: #..#",
		"",
		"...## => #",
		"#.#.# => #",
		"..#.. => .",
	}

	pop := make([]string, 0, 4+(edge*2))
	for i := 0; i < edge; i++ {
		pop = append(pop, ".")
	}
	pop = append(pop, "#")
	pop = append(pop, ".")
	pop = append(pop, ".")
	pop = append(pop, "#")
	for i := 0; i < edge; i++ {
		pop = append(pop, ".")
	}

	must.BeEqual(t, convertInput(input), Input{
		population: pop,
		instructions: []Instruction{
			{".", ".", ".", "#", "#", "#"},
			{"#", ".", "#", ".", "#", "#"},
			{".", ".", "#", ".", ".", "."},
		},
	})
}

// ================== program =============================

type Instruction struct {
	L2 string
	L1 string
	C  string
	R1 string
	R2 string
	N  string
}

func (i Instruction) String() string {
	return i.L2 + i.L1 + i.C + i.R1 + i.R2 + " => " + i.N
}

type Input struct {
	population   []string
	instructions []Instruction
}

func convertInput(input []string) Input {
	pop := strings.TrimPrefix(input[0], "initial state: ")
	res := Input{
		population:   make([]string, 0, len(pop)+(edge*2)),
		instructions: make([]Instruction, 0, len(input)-2),
	}

	for i := 0; i < edge; i++ {
		res.population = append(res.population, ".")
	}
	for _, s := range pop {
		res.population = append(res.population, string(s))
	}
	for i := 0; i < edge; i++ {
		res.population = append(res.population, ".")
	}

	for i := 2; i < len(input); i++ {
		inst := Instruction{
			L2: string(input[i][0]),
			L1: string(input[i][1]),
			C:  string(input[i][2]),
			R1: string(input[i][3]),
			R2: string(input[i][4]),
			N:  string(input[i][9]),
		}
		res.instructions = append(res.instructions, inst)

	}

	return res
}

func grow(input Input, times int) int {

	population := input.population

	for t := 0; t < times; t++ {
		//fmt.Printf("%7d:  %v\n", t, population)
		nextPopulation := make([]string, len(population))
		copy(nextPopulation, population)

		for p := 2; p < len(population)-2; p++ {
			matched := false
			for _, i := range input.instructions {
				if i.L2 == population[p-2] && i.L1 == population[p-1] &&
					i.C == population[p] &&
					i.R1 == population[p+1] && i.R2 == population[p+2] {

					//fmt.Println("instruction matched: ", p, i)
					nextPopulation[p] = i.N
					matched = true
					break
				}
			}

			if !matched {
				nextPopulation[p] = "."
			}
		}

		population = nextPopulation
	}

	fmt.Printf("%7d:  %v\n", times, population)

	sum := 0
	for i, p := range population {
		if p == "#" {
			sum += i - edge
		}
	}
	return sum
}
