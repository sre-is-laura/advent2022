package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	dir      string
	distance int
}

type Point struct {
	x, y int
}

type Board struct {
	tail, head  Point
	tailHistory map[Point]bool
}

func NewBoard() *Board {
	th := make(map[Point]bool)
	b := Board{
		tailHistory: th,
	}
	b.tailHistory[b.tail] = true
	return &b
}

// Main function
func main() {
	//moves := readInput("./testinput.txt")
	moves := readInput("./input.txt")
	//fmt.Printf("Moves %+v\n", moves)

	board := NewBoard()
	for _, m := range moves {
		board.doMove(m)
		//fmt.Printf("Board position after move %+v, tail positions seen %d\n", m, board.countTailPositions())
		//fmt.Printf("Head: %+v, tail %+v\n", board.head, board.tail)
		//board.printTailPositions()
		//fmt.Printf("\n\n")
	}
	fmt.Printf("Tail position count %d\n", board.countTailPositions())
}

func (b *Board) doMove(m Move) {
	// todo
	for i := 0; i < m.distance; i++ {
		b.moveOneStep(m.dir)
	}
}

func (b *Board) moveOneStep(dir string) {
	b.head = getNextPos(dir, b.head)

	// apply planck length rules
	if b.head.x == b.tail.x || b.head.y == b.tail.y { // move if tail is more than 1 behind head
		if b.head.y-b.tail.y > 1 {
			b.tail.y++
		} else if b.tail.y-b.head.y > 1 {
			b.tail.y--
		}

		if b.head.x-b.tail.x > 1 {
			b.tail.x++
		} else if b.tail.x-b.head.x > 1 {
			b.tail.x--
		}
	} else if !touching(b.head, b.tail) { // move diagonally in direction of tail
		if b.head.y > b.tail.y {
			b.tail.y++
		} else {
			b.tail.y--
		}

		if b.head.x > b.tail.x {
			b.tail.x++
		} else {
			b.tail.x--
		}
	}

	b.tailHistory[b.tail] = true
}

func touching(p1 Point, p2 Point) bool {
	return abs(p1.x-p2.x) <= 1 && abs(p1.y-p2.y) <= 1
}

func abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return int(v)
}

func getNextPos(dir string, p Point) Point {
	switch dir {
	case "U":
		return Point{p.x, p.y + 1}
	case "D":
		return Point{p.x, p.y - 1}
	case "R":
		return Point{p.x + 1, p.y}
	case "L":
		return Point{p.x - 1, p.y}
	}

	log.Fatalf("unknown direction")
	return p
}

func (b Board) printTailPositions() {
	maxX := 0
	maxY := 0
	for p, _ := range b.tailHistory {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for i := maxY + 1; i >= 0; i-- {
		for j := 0; j <= maxX+1; j++ {
			p := Point{j, i}
			if b.tail == p {
				fmt.Printf("T")
			} else if b.head == p {
				fmt.Printf("H")
			} else if b.tailHistory[p] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

}

func (b Board) countTailPositions() int {
	result := 0
	for _, _ = range b.tailHistory {
		result++
	}
	return result
}

func readInput(path string) []Move {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := make([]Move, 0)
	for scanner.Scan() {
		val := scanner.Text()
		fields := strings.Split(val, " ")
		dist, _ := strconv.Atoi(fields[1])
		move := Move{fields[0], dist}
		result = append(result, move)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
