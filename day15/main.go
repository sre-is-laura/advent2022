package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Grid struct {
	sbmap map[Point]Point
	minX  int
	maxX  int
	minY  int
	maxY  int
}

func NewGrid(minX int, maxX int, minY int, maxY int, sbmap map[Point]Point) *Grid {
	return &Grid{
		minX:  minX,
		maxX:  maxX,
		minY:  minY,
		maxY:  maxY,
		sbmap: sbmap,
	}
}

// Main function
func main() {
	grid := readInput("./input.txt")

	part1(grid)
	part2(grid)
}

func part1(grid *Grid) {
	row := 2000000 // 10 in test
	c := grid.countPointsWhereBeaconCantExist(row)
	fmt.Printf("Positions where beacon canNOT exist in row %d: %d\n", row, c)
}

func part2(grid *Grid) {
	minxy := 0
	maxxy := 4000000
	p := grid.findBeacon(minxy, maxxy)
	fmt.Printf("Point is %+v, tuning frequency is %d\n", p, p.x*maxxy+p.y)
}

func (g Grid) findBeacon(minxy int, maxxy int) Point {
	for y := minxy; y <= maxxy; y++ {
		x := minxy
		for x <= maxxy {
			p := Point{x, y}
			canBeBeacon := true
			maxNextX := x + 1
			for s, b := range g.sbmap {
				if s.distanceTo(b) >= s.distanceTo(p) {
					canBeBeacon = false
					nxm := x + s.distanceTo(b) - s.distanceTo(p)
					if nxm > maxNextX {
						maxNextX = nxm
					}

				}
			}
			if canBeBeacon {
				return p
			}
			x = maxNextX
		}
	}

	return Point{}
}

func (g Grid) canBeaconExist(p Point) bool {
	for s, b := range g.sbmap {
		if p == b {
			return true // it is a beacon
		}
		if s.distanceTo(b) >= s.distanceTo(p) {
			return false
		}
	}
	return true
}

func (g Grid) countPointsWhereBeaconCantExist(row int) int {
	y := row
	c := 0
	maxdist := 0
	for s, b := range g.sbmap {
		if s.distanceTo(b) > maxdist {
			maxdist = s.distanceTo(b)
		}
	}

	for i := g.minX - maxdist*2; i <= g.maxX+maxdist*2; i++ {
		p := Point{i, y}
		if !g.canBeaconExist(p) {
			c++
		}
	}

	return c

}

func (p Point) distanceTo(o Point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func readInput(path string) *Grid {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	sensors := make([]Point, 0)
	beacons := make([]Point, 0)
	sbmap := make(map[Point]Point)
	minX := 1000000000
	minY := 1000000000
	maxX := 0
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()
		s, b := parseLine(line)
		sbmap[s] = b
		points := []Point{s, b}
		for _, p := range points {
			if p.x < minX {
				minX = p.x
			} else if p.x > maxX {
				maxX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
			if p.y > maxY {
				maxY = p.y
			}
		}
		sensors = append(sensors, s)
		beacons = append(beacons, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid := NewGrid(minX, maxX, minY, maxY, sbmap)

	return grid
}

func parseLine(line string) (Point, Point) {
	fields := strings.Split(line, " ")
	sx, _ := strconv.Atoi(fields[2][2 : len(fields[2])-1])
	sy, _ := strconv.Atoi(fields[3][2 : len(fields[3])-1])
	bx, _ := strconv.Atoi(fields[8][2 : len(fields[8])-1])
	by, _ := strconv.Atoi(fields[9][2:])

	sensor := Point{sx, sy}
	beacon := Point{bx, by}

	return sensor, beacon
}
