package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Grid struct {
	grid []string
}

type Point struct {
	x int
	y int
}

type Path struct {
	path []Point
}

type PathHeap struct {
	paths []*Path
	goal  Point
}

func main() {
	//grid := readInput("./testinput.txt")
	grid := readInput("./input.txt")
	//fmt.Printf("Grid: %+v\n", grid)

	//start := grid.getStart()
	//end := grid.getEnd()
	//fmt.Printf("Start: %+v, end: %v, distance %d\n", start, end, start.distanceTo(end))

	path := grid.findShortestPathFromA()
	path.printPath()
	fmt.Printf("Path: %+v\nSteps %d\n", *path, path.len()-1)
}

func (g Grid) findShortestPathFromA() *Path {
	minPath := g.findPath()

	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid); j++ {
			if g.grid[i][j] == 'a' || g.grid[i][j] == 'S' {
				pt := Point{i, j}
				fmt.Printf("Assessing path from %+v\n", pt)
				path := g.findPathFrom(pt)
				if path.len() > 1 && path.len() < minPath.len() {
					minPath = path
				}
			}
		}
	}
	return minPath
}

func (g Grid) findPath() *Path {
	return g.findPathFrom(g.getStart())
}

func (g Grid) findPathFrom(start Point) *Path {
	openPaths := make([]*Path, 0)
	end := g.getEnd()

	path := NewPath()
	path.addPoint(start)
	openPaths = append(openPaths, path)

	visited := make(map[Point]bool)
	visited[start] = true

	// now iterate towards end
	for len(openPaths) > 0 {
		// expand all current paths to unvisited nodes

		nextOpenPaths := make([]*Path, 0)
		for _, curPath := range openPaths {
			// expand all possible next steps and add to openPaths
			nextPts := g.nextPossiblePoints(curPath)
			for _, p := range nextPts {
				if visited[p] {
					continue
				}
				visited[p] = true
				np := curPath.copy()
				np.addPoint(p)

				//found := p == end
				//fmt.Printf("End is %+v, current point is %+v, found? %v\n", end, p, found)

				if p == end { // Found it
					return np
				} else {
					nextOpenPaths = append(nextOpenPaths, np)
				}
			}
		}
		openPaths = nextOpenPaths
	}
	return path
}

func (g Grid) nextPossiblePoints(path *Path) []Point {
	last := path.last()
	// up, down, l, right
	next := make([]Point, 0)
	next = append(next, Point{last.x, last.y - 1}) // up
	next = append(next, Point{last.x, last.y + 1}) // down
	next = append(next, Point{last.x - 1, last.y}) // left
	next = append(next, Point{last.x + 1, last.y}) // right

	result := make([]Point, 0)
	for _, p := range next {
		if !g.inGrid(p) {
			continue
		}
		if path.contains(p) {
			continue
		}
		if g.allowed(last, p) {
			result = append(result, p)
		}
	}
	return result
}

func (g Grid) allowed(src Point, dst Point) bool {
	srcH := g.grid[src.x][src.y]
	if g.grid[src.x][src.y] == byte('S') {
		srcH = byte('a')
	}

	dstH := g.grid[dst.x][dst.y]
	if g.grid[dst.x][dst.y] == byte('E') {
		dstH = byte('z')
	}
	return dstH <= srcH+1
}

func (g Grid) inGrid(p Point) bool {
	result := p.x >= 0 && p.y >= 0 && p.x < len(g.grid) && p.y < len(g.grid[0])
	return result
}

func (g Grid) getStart() Point {
	return g.getVal('S')
}

func (g Grid) getEnd() Point {
	return g.getVal('E')
}

func (g Grid) getVal(v rune) Point {
	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[i]); j++ {
			if g.grid[i][j] == byte(v) {
				return Point{i, j}
			}
		}
	}
	return Point{-1, -1}
}

func readInput(path string) Grid {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	elevs := make([]string, 0)
	for scanner.Scan() {
		elevs = append(elevs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Grid{elevs}
}

func (p Point) distanceTo(o Point) int {
	// new york block distance
	return abs(o.x-p.x) + abs(o.y-p.y)
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func NewPath() *Path {
	return &Path{
		path: make([]Point, 0),
	}
}

func (p *Path) addPoint(pt Point) {
	p.path = append(p.path, pt)
}

func (p *Path) contains(pt Point) bool {
	for _, p := range p.path {
		if p == pt {
			return true
		}
	}
	return false
}

func (p *Path) len() int {
	return len(p.path)
}

func (p *Path) last() Point {
	return p.path[len(p.path)-1]
}

func (p *Path) copy() *Path {
	cpy := make([]Point, len(p.path))
	copy(cpy, p.path)
	return &Path{cpy}
}

func (p *Path) score(goal Point) int {
	return len(p.path) + p.last().distanceTo(goal)
}

func (p *Path) printPath() {
	h := 0
	w := 0

	pm := make(map[Point]bool)

	for _, pt := range p.path {
		pm[pt] = true
		if pt.x > h {
			h = pt.x
		}
		if pt.y > w {
			w = pt.y
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			pt := Point{i, j}

			if pt == p.last() {
				fmt.Printf("E")
			} else if pt == p.path[0] {
				fmt.Printf("S")
			} else if pm[pt] {
				fmt.Printf("x")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
