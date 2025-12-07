package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func multiply(vals []int) int {
	prod := 1
	for _, v := range vals {
		prod *= v
	}
	return prod
}

func add(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

type problem struct {
	operatorFn func([]int) int
	operands   []int
}

func (p problem) evaluate() int {
	return p.operatorFn(p.operands)
}

func main() {
	rows := loadAndSplitFile("problems.input", "\n")

	total1 := 0
	part1Problems := part1(rows)
	for _, p := range part1Problems {
		total1 += p.evaluate()
	}
	total2 := 0
	problems := part2(rows)
	for _, p := range problems {
		total2 += p.evaluate()
	}
	fmt.Println("total1:", total1)
	fmt.Println("total2:", total2)
}

func InitProblems(rows []string) []problem {
	problems := []problem{}
	parts := strings.Split(strings.TrimSpace(rows[len(rows)-1]), " ")
	for _, part := range parts {
		if part == "" {
			continue
		}
		var operatorFn func([]int) int
		switch part {
		case "*":
			operatorFn = multiply
		case "+":
			operatorFn = add
		}
		problems = append(problems, problem{operatorFn: operatorFn, operands: []int{}})
	}
	return problems
}

func part1(rows []string) []problem {
	problems := InitProblems(rows)
	for _, row := range rows[:len(rows)-1] {
		parts := strings.Split(strings.TrimSpace(row), " ")
		index := 0
		for _, part := range parts {
			if part == "" {
				continue
			}
			val, _ := strconv.Atoi(part)
			problems[index].operands = append(problems[index].operands, val)
			index++
		}
	}
	return problems
}

func part2(rows []string) []problem {
	problems := InitProblems(rows)
	numbersStrs := make([]string, len(([]rune)(rows[0])))
	for _, row := range rows[:len(rows)-1] {
		parts := []rune(row)
		for i, part := range parts {
			numbersStrs[i] += string(part)
		}
	}
	problemCounter := 0
	for _, numbersStr := range numbersStrs {
		if strings.TrimSpace(numbersStr) == "" {
			problemCounter++
			continue
		}
		val, _ := strconv.Atoi(strings.TrimSpace(numbersStr))
		problems[problemCounter].operands = append(problems[problemCounter].operands, val)

	}
	return problems
}

func loadAndSplitFile(path string, sep string) []string {
	dat, _ := os.ReadFile(path)
	code := string(dat)
	return strings.Split(code, sep)
}

func parseRange(s string) (int, int) {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}
