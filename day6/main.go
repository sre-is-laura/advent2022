package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Main function
func main() {
	runTests()

	input := readInput("./input.txt")
	//result := charsToPacketStartMarker(input) // part1
	result := charsToMessageStartMarker(input) // part1
	fmt.Printf("Result is %d\n", result)
}

func charsToPacketStartMarker(input string) int {
	return charsToUniqueSeq(input, 4)
}

func charsToMessageStartMarker(input string) int {
	return charsToUniqueSeq(input, 14)
}

func charsToUniqueSeq(input string, seqlen int) int {
	for i := seqlen; i < len(input); i++ {
		precSeq := input[i-seqlen : i]
		if allDifferent((precSeq)) {
			return i
		}
	}

	log.Fatalf("No marker found for %s\n", input)
	return 0
}

func allDifferent(input string) bool {
	seen := make(map[rune]bool)

	for _, r := range input {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func runTests() {
	type test struct {
		input        string
		packetStart  int
		messageStart int
	}

	tests := []test{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", packetStart: 7, messageStart: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", packetStart: 5, messageStart: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", packetStart: 6, messageStart: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", packetStart: 10, messageStart: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", packetStart: 11, messageStart: 26},
	}

	for _, tc := range tests {
		gotPs := charsToPacketStartMarker(tc.input)
		gotMs := charsToMessageStartMarker(tc.input)
		if tc.packetStart != gotPs {
			fmt.Printf("Test failed to get packet start marker: %s want %d got %d\n", tc.input, tc.packetStart, gotPs)
		}
		if tc.messageStart != gotMs {
			fmt.Printf("Test failed to get message start marker: %s want %d got %d\n", tc.input, tc.messageStart, gotMs)
		}
	}

}

func readInput(path string) string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	val := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return val
}
