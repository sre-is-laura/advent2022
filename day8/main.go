package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Main function
func main() {
	//trees := readInput("./testinput.txt")
	trees := readInput("./input.txt")
	//fmt.Printf("Trees: %+v\n", trees)

	//count := countVisibleTrees(trees)
	//fmt.Printf("Visible trees: %d\n", count)

	scenic := maxScenicScore(trees)
	fmt.Printf("Best scenic score %d\n", scenic)

}

func countVisibleTrees(trees [][]int) int {
	w := len(trees[0])
	h := len(trees)

	countVisible := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if visible(trees, i, j) {
				countVisible++
			}
		}
	}

	return countVisible
}

func maxScenicScore(trees [][]int) int {
	w := len(trees[0])
	h := len(trees)

	max := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			score := scenicScore(trees, i, j)
			if score > max {
				max = score
			}
		}
	}

	return max
}

func scenicScore(trees [][]int, x int, y int) int {
	th := trees[x][y]
	w := len(trees[0])
	h := len(trees)

	if x == 0 || y == 0 || x == h-1 || y == w-1 {
		return 0
	}

	up := 0
	for i := x - 1; i >= 0; i-- {
		up++
		if trees[i][y] >= th {
			break
		}
	}

	down := 0
	for i := x + 1; i < h; i++ {
		down++
		if trees[i][y] >= th {
			break
		}
	}

	left := 0
	for j := y - 1; j >= 0; j-- {
		left++
		if trees[x][j] >= th {
			break
		}
	}

	right := 0
	for j := y + 1; j < w; j++ {
		right++
		if trees[x][j] >= th {
			break
		}
	}
	score := up * down * left * right
	return score
}

// Is tree visible from edge
func visible(trees [][]int, x int, y int) bool {
	th := trees[x][y]
	w := len(trees[0])
	h := len(trees)

	if x == 0 || y == 0 || x == h-1 || y == w-1 {
		return true
	}

	maybeVisible := true
	for i := 0; i < x; i++ {
		if trees[i][y] >= th {
			maybeVisible = false
			break
		}
	}
	if maybeVisible {
		return true
	}

	maybeVisible = true
	for i := x + 1; i < h; i++ {
		if trees[i][y] >= th {
			maybeVisible = false
			break
		}
	}
	if maybeVisible {
		return true
	}

	maybeVisible = true
	for j := 0; j < y; j++ {
		if trees[x][j] >= th {
			maybeVisible = false
			break
		}
	}
	if maybeVisible {
		return true
	}

	maybeVisible = true
	for j := y + 1; j < w; j++ {
		if trees[x][j] >= th {
			maybeVisible = false
			break
		}
	}

	return maybeVisible
}

func readInput(path string) [][]int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	input := make([]string, 0)
	for scanner.Scan() {
		val := scanner.Text()
		input = append(input, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := make([][]int, len(input))
	for i := 0; i < len(input[0]); i++ {
		result[i] = make([]int, len(input[0]))
		chars := strings.Split(input[i], "")
		for j, c := range chars {
			h, _ := strconv.Atoi((c))
			result[i][j] = h
		}
	}

	return result
}
