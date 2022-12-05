package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Vals []string
}

type Move struct {
	Count   int
	Fromidx int
	Toidx   int
}

func NewStack() *Stack {
	s := Stack{}
	s.Vals = make([]string, 0)
	return &s
}

// The stack is 'upside-down', i.e. 0 is the top
// this function appends to the end, i.e. the bottom. Only used for parsing
func (s *Stack) Append(item string) {
	s.Vals = append(s.Vals, item)
}

func (s *Stack) Push(items []string) {
	ns := make([]string, 0)
	ns = append(ns, items...)

	s.Vals = append(ns, s.Vals...)
}

func (s *Stack) Top() string {
	if len(s.Vals) > 0 {
		val := s.Vals[0]
		return val
	} else {
		return ""
	}
}

func (s *Stack) Popitems(count int) []string {
	val := s.Vals[0:count]
	s.Vals = s.Vals[count:len(s.Vals)]
	return val
}

// Main function
func main() {
	//stacks, moves := readStacksAndMoves("./testinput.txt")
	stacks, moves := readStacksAndMoves("./input.txt")
	//PrintStacks(stacks)

	//part1(stacks, moves)
	part2(stacks, moves)

	result := ""
	for _, s := range stacks {
		//fmt.Printf("Stack: %+v Top %s\n", s, s.Top())
		result += s.Top()
	}
	fmt.Printf("Result %s\n", result)
}

func part2(stacks []*Stack, moves []Move) {
	for _, m := range moves {
		items := stacks[m.Fromidx-1].Popitems(m.Count)
		stacks[m.Toidx-1].Push(items)
		//PrintStacks(stacks)
	}
}

func part1(stacks []*Stack, moves []Move) {
	for _, m := range moves {
		for i := 0; i < m.Count; i++ {
			//fmt.Printf("Move: %+v\n", m)
			item := stacks[m.Fromidx-1].Popitems(1)
			stacks[m.Toidx-1].Push(item)
			//PrintStacks(stacks)
		}
	}
}

func PrintStacks(stacks []*Stack) {
	for _, s := range stacks {
		fmt.Printf("Stack: %+v\n", s)
	}
}

func readStacksAndMoves(path string) ([]*Stack, []Move) {
	stacks := make([]*Stack, 0)
	moves := make([]Move, 0)

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	readingMoves := false
	for scanner.Scan() {
		val := scanner.Text()

		if !readingMoves && strings.Count(val, "[") == 0 {
			readingMoves = true
			scanner.Scan()
			continue
		} else if readingMoves {
			fields := strings.Split(val, " ")
			m := Move{}
			m.Count, _ = strconv.Atoi(fields[1])
			m.Fromidx, _ = strconv.Atoi(fields[3])
			m.Toidx, _ = strconv.Atoi(fields[5])
			moves = append(moves, m)
		} else {
			for i := 0; i <= len(val)/4; i++ {
				if val[i*4:i*4+1] == "[" {
					item := val[i*4+1 : i*4+2]
					for len(stacks) < i+1 {
						stacks = append(stacks, NewStack())
					}
					stacks[i].Append(item)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stacks, moves
}
