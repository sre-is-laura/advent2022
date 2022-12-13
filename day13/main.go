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

type PacketContentType int

type Packet struct {
	content string
}

type PacketPair struct {
	left  Packet
	right Packet
}

type PacketContent struct {
	val    int
	hasval bool
	pc     []PacketContent
}

// Main function
func main() {
	//pairs := readInput("./testinput.txt")
	pairs := readInput("./input.txt")
	//fmt.Printf("Packet pairs: %+v\n", pairs)

	part1(pairs)
	part2(pairs)
}

func part2(pairs []PacketPair) {
	packets := make([]Packet, 0)

	div1 := Packet{"[[2]]"}
	div2 := Packet{"[[6]]"}

	packets = append(packets, div1)
	packets = append(packets, div2)

	for _, p := range pairs {
		packets = append(packets, p.left)
		packets = append(packets, p.right)

	}

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].getContent().Compare(packets[j].getContent()) == -1
	})

	//fmt.Printf("Part2: All packets\n")
	//for _, p := range packets {
	//fmt.Printf("%s\n", p.content)
	//}

	div1ind := 0
	div2ind := 0
	for i, p := range packets {
		if p == div1 {
			div1ind = i + 1
		}
		if p == div2 {
			div2ind = i + 1
		}
	}
	fmt.Printf("Part2: result is %d\n", div1ind*div2ind)
}

func part1(pairs []PacketPair) {
	sumidx := 0
	for i, p := range pairs {
		if p.inOrder() {
			sumidx += i + 1
		}
	}
	fmt.Printf("Part1: Sum of indexes is %d\n", sumidx)
}

func (p PacketPair) inOrder() bool {
	return p.left.getContent().Compare(p.right.getContent()) <= 0
}

// 1 if left greater than right
// 0 if equal
// -1 if left less than right
func (left PacketContent) Compare(right PacketContent) int {
	//fmt.Printf("Comparing %s and %s\n", left.String(), right.String())
	//fmt.Printf("Comparing %+v and %+v\n", left, right)

	didRightConversion := false
	if left.hasval && right.hasval {
		if left.val > right.val {
			return 1
		} else if left.val == right.val {
			return 0
		} else {
			return -1
		}
	} else if left.hasval { // left an int, right a list
		left = PacketContent{pc: []PacketContent{left}}
	} else if right.hasval { // left a nonempty list, right an int
		right = PacketContent{pc: []PacketContent{right}}
		didRightConversion = true
	}

	// both lists, iterate, return true if right runs out first
	for i := 0; i < len(left.pc); i++ {
		if i >= len(right.pc) {
			if didRightConversion {
				return 1
			}
			return 1 // left > right
		} else if left.pc[i].Compare(right.pc[i]) == 1 {
			return 1 // left > right
		} else if left.pc[i].Compare(right.pc[i]) == -1 {
			return -1 // left < right
		}
	}
	if len(left.pc) < len(right.pc) {
		return -1
	} else {
		return 0
	}
}

func (pc PacketContent) String() string {
	result := ""
	if pc.hasval {
		result += fmt.Sprintf("%d", pc.val)
	} else {
		result += "["
		for i, c := range pc.pc {
			result += c.String()
			if i < len(pc.pc)-1 {
				result += ","
			}
		}
		result += "]"
	}
	return result
}

func (p Packet) getContent() PacketContent {
	content := getContentRec(p.content)
	//fmt.Printf("%s -> %+v\n", p.content, content)
	return PacketContent{pc: content}
}

func getContentRec(input string) []PacketContent {
	//fmt.Printf("GetContentRec called: %s\n", input)
	result := make([]PacketContent, 0)
	if len(input) == 0 {
		return result
	}

	// consume first item - should be int or start of a list
	if input[0] == '[' { // list start
		listEnd := findListEnd(input)

		ic := getContentRec(input[1:listEnd])
		pc := PacketContent{pc: ic}
		result = append(result, pc)

		if listEnd >= len(input)-1 {
			return result
		} else {
			return append(result, getContentRec(input[listEnd+2:len(input)])...)
		}
	} else {
		// expect an int and a comma, consume and recurse

		comma := strings.Index(input, ",")

		val := input
		if comma != -1 {
			val = input[0:comma]
		}
		iv, _ := strconv.Atoi(val)

		pc := PacketContent{val: iv, hasval: true}
		result = append(result, pc)

		if comma != -1 {
			return append(result, getContentRec(input[comma+1:len(input)])...)
		} else {
			return result
		}
	}

	log.Fatalf("should not be here")
	return make([]PacketContent, 0)
}

// for the [ at index 0, return index of matching
func findListEnd(input string) int {
	openCount := 1

	for i := 1; i < len(input); i++ {
		if input[i] == '[' {
			openCount++
		} else if input[i] == ']' {
			openCount--
		}
		if openCount == 0 {
			return i
		}
	}

	log.Fatalf("should not be here")
	return -1
}

func readInput(path string) []PacketPair {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)
	for scanner.Scan() {
		val := scanner.Text()
		lines = append(lines, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := make([]PacketPair, 0)
	for i := 0; i < len(lines); i += 3 {
		f := Packet{lines[i]}
		s := Packet{lines[i+1]}
		result = append(result, PacketPair{f, s})
	}

	return result
}
