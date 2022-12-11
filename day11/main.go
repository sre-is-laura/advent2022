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

type Monkey struct {
	items        []uint64
	inspectCount int
	op           Operation
	test         Test
}

func NewMonkey() *Monkey {
	return &Monkey{
		items: make([]uint64, 0),
	}
}

type Operation struct {
	right    string
	operator string
}

type Test struct {
	divisor   uint64
	truePath  int
	falsePath int
}

func main() {
	//monkeys := readInput("./testinput.txt")
	monkeys := readInput("./input.txt")
	fmt.Printf("Monkeys: %+v\n", monkeys)

	// product of all the divisors can be used to prevent runaway growth of item worry scores
	divs := uint64(1)
	for _, m := range monkeys {
		divs *= m.test.divisor
	}
	fmt.Printf("Common divisor is %d\n", divs)

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.doRound(monkeys, divs, 1)
		}

		if i%1000 == 0 {
			for j, m := range monkeys {
				fmt.Printf("After %d rounds monkey %d has inspect count %d\n", i, j, m.inspectCount)
			}
		}
	}

	ics := make([]int, len(monkeys))
	for i, m := range monkeys {
		ics[i] = m.inspectCount
	}
	sort.Ints(ics)

	fmt.Printf("Monkey business score: %d\n", ics[len(ics)-1]*ics[len(ics)-2])
}

func (m *Monkey) doRound(monkeys []*Monkey, commonDiv uint64, worryDivisor uint64) {
	for _, item := range m.items {
		m.inspectAndThrowItem(monkeys, item, commonDiv, worryDivisor)
	}
	m.items = make([]uint64, 0)
}

func (m *Monkey) inspectAndThrowItem(monkeys []*Monkey, item uint64, commonDiv uint64, worryDivisor uint64) {
	worryLevel := m.op.calculate(item)
	worryLevel /= worryDivisor

	// only the modulo of product of all divisors affects outcome, so let's not overflow
	if worryLevel >= commonDiv {
		worryLevel = worryLevel % commonDiv
	}

	fmt.Printf("New worry level is %d\n", worryLevel)

	m.inspectCount++

	nextMonkey := m.test.run(worryLevel)
	monkeys[nextMonkey].catchItem(worryLevel)
	//fmt.Printf("Item worry value: %d thrown to monkey %d\n", worryLevel, nextMonkey)
}

func (m *Monkey) catchItem(item uint64) {
	m.items = append(m.items, item)
}

func (t Test) run(item uint64) int {
	if item%t.divisor == 0 {
		return t.truePath
	} else {
		return t.falsePath
	}
}

func (op Operation) calculate(item uint64) uint64 {
	right := uint64(0)
	if op.right == "old" {
		right = item
	} else {
		var err error
		r, err := strconv.Atoi(op.right)
		right = uint64(r)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	if op.operator == "*" {
		return item * right
	} else if op.operator == "+" {
		return item + right
	}

	log.Fatalf("unknown operation %s", op.operator)
	return 0
}

func readInput(path string) []*Monkey {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	result := make([]*Monkey, 0)
	for scanner.Scan() {
		val := scanner.Text()

		if len(val) == 0 {
			continue
		} else {
			if strings.HasPrefix(val, "Monkey") {
				m := NewMonkey()
				scanner.Scan()

				items := scanner.Text()[18:]
				fs := strings.Split(items, ", ")
				for _, f := range fs {
					i, _ := strconv.Atoi(f)
					m.items = append(m.items, uint64(i))
				}

				scanner.Scan()
				optext := scanner.Text()[23:]
				fs = strings.Split(optext, " ")
				m.op.operator = fs[0]
				m.op.right = fs[1]

				scanner.Scan()
				div := scanner.Text()[21:]
				d, _ := strconv.Atoi(div)
				m.test.divisor = uint64(d)
				scanner.Scan()

				tp := scanner.Text()[29:]
				m.test.truePath, _ = strconv.Atoi(tp)

				scanner.Scan()
				fp := scanner.Text()[30:]
				m.test.falsePath, _ = strconv.Atoi(fp)

				//fmt.Printf("New monkey %+v\n", *m)
				result = append(result, m)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
