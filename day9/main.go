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
	knots       []Point
	tailHistory map[Point]bool
}

func NewBoard(numTails int) *Board {
	th := make(map[Point]bool)
	knots := make([]Point, numTails+1)
	b := Board{
		tailHistory: th,
		knots:       knots,
	}
	b.tailHistory[b.tail()] = true
	return &b
}

// Main function
func main() {
	//moves := readInput("./testinput2.txt")
	moves := readInput("./input.txt")
	//fmt.Printf("Moves %+v\n", moves)

	board := NewBoard(9)
	for _, m := range moves {
		board.doMove(m)
		/*fmt.Printf("Board position after move %+v, tail positions seen %d\n", m, board.countTailPositions())
		fmt.Printf("Head: %+v, tail %+v\n", board.head(), board.tail())
		board.printTailPositions()
		fmt.Printf("\n\n")*/
	}
	fmt.Printf("Tail position count %d\n", board.countTailPositions())
}

func (b *Board) doMove(m Move) {
	for i := 0; i < m.distance; i++ {
		b.moveOneStep(m.dir)
	}
}

func (b *Board) moveOneStep(dir string) {
	b.knots[0] = getNextPos(dir, b.head())

	for i := 1; i < len(b.knots); i++ {
		b.knots[i] = planckRules(b.knots[i-1], b.knots[i])
	}

	b.tailHistory[b.tail()] = true
}

func planckRules(head Point, tail Point) Point {
	// apply planck length rules
	if head.x == tail.x || head.y == tail.y { // move if tail is more than 1 behind head
		if head.y-tail.y > 1 {
			tail.y++
		} else if tail.y-head.y > 1 {
			tail.y--
		}

		if head.x-tail.x > 1 {
			tail.x++
		} else if tail.x-head.x > 1 {
			tail.x--
		}
	} else if !touching(head, tail) { // move diagonally in direction of tail
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}

		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
	}
	return tail
}

func (b *Board) tail() Point {
	return b.knots[len(b.knots)-1]
}

func (b *Board) head() Point {
	return b.knots[0]
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
			if b.tail() == p {
				fmt.Printf("T")
			} else if b.head() == p {
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
