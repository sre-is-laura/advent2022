package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ElfPair struct {
	elf1 Range
	elf2 Range
}

type Range struct {
	start int
	end   int
}

// Main function
func main() {
	elfpairs := readElfPairs("./input.txt")
	//elfpairs := readElfPairs("./testinput.txt")
	fmt.Printf("ElfPairs: %+v\n", elfpairs)

	contains := 0
	overlaps := 0
	for _, ep := range elfpairs {
		if ep.fullyContains() {
			contains++
			//fmt.Printf("Fully contained: %+v\n", ep)
		}

		if ep.overlaps() {
			fmt.Printf("Has overlaps: %+v\n", ep)
			overlaps++
		}
	}
	fmt.Printf("Pairs fully contained: %d\n", contains)
	fmt.Printf("Pairs with overlaps: %d\n", overlaps)
}

func (ep ElfPair) fullyContains() bool {
	return ep.elf1.contains(ep.elf2) || ep.elf2.contains(ep.elf1)
}

func (ep ElfPair) overlaps() bool {
	return ep.elf1.overlaps(ep.elf2)
}

func (r Range) contains(or Range) bool {
	return r.start <= or.start && r.end >= or.end
}
func (r Range) overlaps(or Range) bool {
	return (r.end >= or.start && r.start <= or.start) || (or.end >= r.start && or.start <= r.start)
}

func readElfPairs(path string) []ElfPair {
	result := make([]ElfPair, 0)

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		pairs := strings.Split(val, ",")
		e1r := getRange(pairs[0])
		e2r := getRange(pairs[1])
		ep := ElfPair{e1r, e2r}

		result = append(result, ep)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func getRange(str string) Range {
	p := strings.Split(str, "-")
	ps, _ := strconv.Atoi(p[0])
	pe, _ := strconv.Atoi(p[1])
	return Range{ps, pe}
}
