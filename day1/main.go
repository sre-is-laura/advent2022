package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Main function
func main() {
	cals := readCalories("./input.txt")
	//fmt.Printf("Cals: %v\n", cals)

	topThree := 0
	sort.Ints(cals)
	for i := len(cals) - 3; i < len(cals); i++ {
		topThree += cals[i]
	}
	fmt.Printf("Calories carried by top three elves %d\n", topThree)
}

func part1(cals []int) {
	maxElf := 0

	for i := 1; i < len(cals); i++ {
		if cals[i] > cals[maxElf] {
			maxElf = i
		}
	}

	fmt.Printf("Max elf: %d; Calories carried %d\n", maxElf+1, cals[maxElf]) // zero offset
}

func readCalories(path string) []int {
	result := make([]int, 0)

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	currentCals := 0
	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		if len(val) == 0 {
			result = append(result, currentCals)
			currentCals = 0
		} else {
			i, err := strconv.Atoi(val)
			currentCals += i
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	result = append(result, currentCals)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
