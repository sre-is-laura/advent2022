package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Main function
func main() {
	rounds := readRounds("./input.txt")

	fmt.Printf("Rounds %+v %d\n", rounds, len(rounds))
	score := 0

	for _, r := range rounds {
		score += r.Score()
		fmt.Printf("Rounds %v score %d\n", r, r.Score())
	}
	fmt.Printf("Score: %d\n", score)
}

const (
	Rock     string = "X"
	Paper    string = "Y"
	Scissors string = "Z"
)

const (
	OppRock     string = "A"
	OppPaper    string = "B"
	OppScissors string = "C"
)

const (
	Win  int = 6
	Lose int = 0
	Draw int = 3
)

type Round struct {
	oppMove string
	move    string
}

func (r Round) Score() int {
	switch r.oppMove {
	case OppRock:
		if r.move == Rock {
			return Draw + val(r.move)
		} else if r.move == Paper {
			return Win + val(r.move)
		} else {
			return Lose + val(r.move)
		}
	case OppPaper:
		if r.move == Rock {
			return Lose + val(r.move)
		} else if r.move == Paper {
			return Draw + val(r.move)
		} else {
			return Win + val(r.move)
		}
	case OppScissors:
		if r.move == Rock {
			return Win + val(r.move)
		} else if r.move == Paper {
			return Lose + val(r.move)
		} else {
			return Draw + val(r.move)
		}
	}
	log.Fatal("wat")
	return 0
}

func val(move string) int {
	switch move {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	log.Fatal("wat")
	return 0
}

func readRounds(path string) []Round {
	result := make([]Round, 0)

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		fields := strings.Split(val, " ")

		if len(fields) != 2 {
			log.Fatal("can't parse")
		} else {
			r := Round{fields[0], fields[1]}
			result = append(result, r)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
