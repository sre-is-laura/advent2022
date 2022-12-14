package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type GridState int

const (
	Empty GridState = iota
	Rock
	Sand
	SandOrigin
	Floor
)

type Point struct {
	x int
	y int
}

type Grid struct {
	points [][]GridState
	minX   int
	maxX   int
	minY   int
	maxY   int
}

func NewGrid(minX int, maxX int, minY int, maxY int) *Grid {
	height := maxY - minY + 3
	width := maxX - minX + height*2

	points := make([][]GridState, height)
	for i := 0; i < height; i++ {
		points[i] = make([]GridState, width)
	}
	for j := 0; j < width; j++ {
		points[height-1][j] = Floor
	}

	return &Grid{
		minX:   minX - height,
		maxX:   maxX + height,
		minY:   minY,
		maxY:   maxY + 2,
		points: points,
	}

}

func GetSandOrigin() Point {
	return Point{500, 0}
}

// Main function
func main() {
	//grid := readInput("./testinput.txt")
	grid := readInput("./input.txt")

	result := 0
	for grid.ProduceSand() {
		result++
	}
	fmt.Printf("Part 2: can produce %d units of sand\n", result+1)

	//grid.Print()
}

// returns false if comes to rest at SandOrigin
func (g *Grid) ProduceSand() bool {
	// sand comes from sand origin
	so := GetSandOrigin()
	return g.settleSand(so) != so
}

func (g *Grid) settleSand(sp Point) Point {
	// at this point there is something underneath
	// moves diagonally left if not blocked
	under := Point{sp.x, sp.y + 1}
	diagLeft := Point{sp.x - 1, sp.y + 1}
	diagRight := Point{sp.x + 1, sp.y + 1}

	if g.GetPoint(under) == Floor {
		g.SetPoint(sp, Sand)
		return sp
	}
	if g.GetPoint(under) == Empty {
		// move into it, recurse
		return g.settleSand(under)
	}
	if g.GetPoint(diagLeft) == Floor {
		g.SetPoint(sp, Sand)
		return sp
	}
	if g.GetPoint(diagLeft) == Empty {
		// move into it, recurse
		return g.settleSand(diagLeft)
	}

	if g.GetPoint(diagRight) == Floor {
		g.SetPoint(sp, Sand)
		return sp
	}
	if g.GetPoint(diagRight) == Empty {
		return g.settleSand(diagRight)
	}
	// sand is settled
	g.SetPoint(sp, Sand)
	return sp
}

func (g *Grid) GetPoint(pt Point) GridState {
	if pt.x < g.minX || pt.x > g.maxX || pt.y < g.minY || pt.y > g.maxY {
		return Floor
	}

	return g.points[pt.y-g.minY][pt.x-g.minX]
}

func (g *Grid) SetPoint(pt Point, val GridState) {
	g.points[pt.y-g.minY][pt.x-g.minX] = val
}

func (g *Grid) DrawRockLine(pt1 Point, pt2 Point) {
	if pt1.x == pt2.x {
		incr := 1
		if pt1.y > pt2.y {
			incr = -1
		}
		for i := pt1.y; i != pt2.y+incr; i += incr {
			pt := Point{pt1.x, i}
			g.SetPoint(pt, Rock)
		}
	} else if pt1.y == pt2.y {
		incr := 1
		if pt1.x > pt2.x {
			incr = -1
		}
		for i := pt1.x; i != pt2.x+incr; i += incr {
			pt := Point{i, pt1.y}
			g.SetPoint(pt, Rock)
		}
	}
	// todo lines the other direction
}

func (g Grid) Print() {
	fmt.Println("\n")
	for _, line := range g.points {
		for _, pt := range line {
			if pt == Empty {
				fmt.Printf(".")
			} else if pt == Rock {
				fmt.Printf("#")
			} else if pt == Sand {
				fmt.Printf("o")
			} else if pt == SandOrigin {
				fmt.Printf("+")
			} else if pt == Floor {
				fmt.Printf("_")
			}
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

func readInput(path string) *Grid {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	paths := make([][]Point, 0)
	minX := 1000000000
	maxX := 0
	maxY := 0

	for scanner.Scan() {
		val := scanner.Text()
		points := strings.Split(val, " -> ")

		path := make([]Point, 0)
		for _, p := range points {
			crds := strings.Split(p, ",")
			x, _ := strconv.Atoi(crds[0])
			y, _ := strconv.Atoi(crds[1])

			pt := Point{x, y}
			path = append(path, pt)
			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
		paths = append(paths, path)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid := NewGrid(minX, maxX, 0, maxY)
	//fmt.Printf("Grid: %+v\n\n", *grid)

	grid.SetPoint(GetSandOrigin(), SandOrigin)
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			grid.DrawRockLine(path[i], path[i+1])
		}
	}

	return grid
}
