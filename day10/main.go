package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	cmd   string
	param int
}

type SimpleCPU struct {
	signalStrengths map[int]int
	xReg            int
	cycleCount      int
	crt             []string
}

func NewSimpleCPU() *SimpleCPU {
	ss := make(map[int]int)
	crt := make([]string, 6)
	cpu := SimpleCPU{
		signalStrengths: ss,
		xReg:            1,
		crt:             crt,
	}
	return &cpu
}

// Main function
func main() {
	//instrs := readInput("./testinput.txt")
	instrs := readInput("./input.txt")

	cpu := NewSimpleCPU()
	for _, i := range instrs {
		cpu.exec(i)
	}

	//part1(cpu)
	cpu.printCRT()
}

func part1(cpu *SimpleCPU) {
	sum := cpu.getSignalStrength(20) + cpu.getSignalStrength(60) + cpu.getSignalStrength(100) + cpu.getSignalStrength(140) + cpu.getSignalStrength(180) + cpu.getSignalStrength(220)
	fmt.Printf("Signal strengths sum %d\n", sum)
}

func (c *SimpleCPU) printCRT() {
	for _, s := range c.crt {
		fmt.Printf("%s\n", s)
	}
}

func (c *SimpleCPU) exec(instr Instruction) {

	switch instr.cmd {
	case "noop":
		// do nothing
		c.incrCycleCount()
	case "addx":
		c.incrCycleCount()
		c.incrCycleCount()
		c.xReg += instr.param
	}
}

func (c *SimpleCPU) incrCycleCount() {
	c.drawPixel()
	c.cycleCount++
	if c.cycleCount == 20 || (c.cycleCount-20)%40 == 0 {
		c.signalStrengths[c.cycleCount] = c.xReg * c.cycleCount
	}
}

func (c *SimpleCPU) drawPixel() {
	drawCol := (c.cycleCount) % 40
	drawRow := int((c.cycleCount - drawCol) / 40)

	if c.spritePos(drawCol) {
		c.crt[drawRow] += "#"
	} else {
		c.crt[drawRow] += " "
	}
}

func (c *SimpleCPU) spritePos(x int) bool {
	result := x >= c.xReg-1 && x <= c.xReg+1
	return result
}

func (c *SimpleCPU) getSignalStrength(i int) int {
	return c.signalStrengths[i]
}

func readInput(path string) []Instruction {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := make([]Instruction, 0)
	for scanner.Scan() {
		val := scanner.Text()

		fs := strings.Split(val, " ")
		i := Instruction{cmd: fs[0]}
		if len(fs) == 2 {
			p, _ := strconv.Atoi(fs[1])
			i.param = p
		} else if len(fs) > 2 {
			log.Fatalf("can't parse %s\n", val)
		}
		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
