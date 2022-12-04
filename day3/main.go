package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Backpack struct {
	Compartment1 []rune
	Compartment2 []rune
}

// rune for a is 97, z is 122
// rune for A is 65, z is 90

// Main function
func main() {
	//backpacks := readBackpacks("./testinput.txt")
	backpacks := readBackpacks("./input.txt")
	//fmt.Printf("Backpacks: %+v\n", backpacks)

	part2(backpacks)
}

func part2(backpacks []Backpack) {
	sumP := 0

	for i := 0; i < len(backpacks)-2; i += 3 {
		c := commonItem(backpacks[i : i+3])
		p := priority(c)

		//fmt.Printf("Common item for backpacks %vis %v/%s with priority %d\n", backpacks[i:i+3], c, string(c), p)

		sumP += p
	}
	fmt.Printf("Sum of priorities of common items: %d\n", sumP)
}

// 2770 is too high
func part1(backpacks []Backpack) {
	sumP := 0
	for _, bp := range backpacks {
		r := bp.getCommonItem()
		p := priority(r)
		sumP += p
	}
	fmt.Printf("Sum of priorities: %d\n", sumP)
}

func commonItem(bps []Backpack) rune {
	maps := make([]map[rune]bool, 0)

	for _, bp := range bps {
		maps = append(maps, bp.getItemMap())
	}

	keys := make([]rune, 0, len(maps[0]))
	for k := range maps[0] {
		keys = append(keys, k)
	}

	for _, k := range keys {

		maybeFound := true
		for _, m := range maps {
			if m[k] != true {
				maybeFound = false
				continue
			}
		}

		if maybeFound {
			return k
		}
	}

	log.Fatal("whoops can't find a common item")
	return 'a'
}

func priority(r rune) int {
	i := int(r)

	if i <= 90 && i >= 65 {
		return i - 65 + 27
	} else if i <= 122 && i >= 97 {
		return i - 96
	}

	fmt.Printf("Can't compute priority for rune : %v\n", r)

	log.Fatal("whoops")
	return 0
}

func (bp Backpack) getItemMap() map[rune]bool {
	itemMap := make(map[rune]bool)
	for _, r := range bp.Compartment1 {
		itemMap[r] = true
	}
	for _, r := range bp.Compartment2 {
		itemMap[r] = true
	}
	return itemMap
}

func (bp Backpack) getCommonItem() rune {
	for _, i := range bp.Compartment1 {
		for _, j := range bp.Compartment2 {
			if i == j {
				return i
			}
		}
	}
	log.Fatal("whoops")
	return []rune("a")[0]
}

func readBackpacks(path string) []Backpack {
	result := make([]Backpack, 0)

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		runes := []rune(val)
		bp := Backpack{
			Compartment1: runes[0 : len(runes)/2],
			Compartment2: runes[len(runes)/2:],
		}

		result = append(result, bp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
