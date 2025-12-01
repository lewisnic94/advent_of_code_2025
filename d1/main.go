package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg1 := os.Args[1]
	path := "input.txt"
	dat, _ := os.ReadFile(path)

	code := string(dat)
	items := strings.Split(code, "\n")
	commands := make([]command, 0, len(items))
	for _, item := range items {
		if item != "" {
			commands = append(commands, command(item))
		}
	}

	startingPoint := 50
	switch arg1 {
	case "algo1":
		finalPosition := decrypt1(startingPoint, commands)
		fmt.Println("Final Position:", finalPosition)
		return
	case "algo2":
		finalPosition := decrypt2(startingPoint, commands)
		fmt.Println("Final Position:", finalPosition)
		return
	}
}

type command string

func (c command) direction() string {
	return string(c[0])
}

func (c command) distance() int {
	distanceStr := strings.Trim(string(c), "LR")
	distance, _ := strconv.Atoi(distanceStr)
	return distance
}

// when we we hit AFTER finishing a command modulo 100 == 0, we increment a counter
func decrypt1(start int, commands []command) int {
	counter := 0
	for _, cmd := range commands {
		switch cmd.direction() {
		case "R":
			start += cmd.distance()
		case "L":
			start -= cmd.distance()
		}
		fmt.Printf("After command %s, position is %d\n", cmd, start)
		if start%100 == 0 {
			counter += 1
		}
	}
	return counter
}

// In the process of moving, if we PASS through a position that is a multiple of 100,
// we count that as a special event. Return the count of such events.
func decrypt2(start int, commands []command) int {
	counter := 0
	for _, cmd := range commands {
		travel := cmd.distance()
		var d2cL, d2cR int
		if start >= 0 {
			d2cL = start % 100
			d2cR = 100 - d2cL
			if d2cL == 0 {
				d2cL = 100
				d2cR = 100
			}
		} else {
			// for negative numbers, we need to adjust the calculation
			d2cR = int(math.Abs(float64(start % -100)))
			d2cL = 100 - d2cR
			if d2cR == 0 {
				d2cR = 100
				d2cL = 100
			}
		}
		var end int
		var distance int
		switch cmd.direction() {
		case "R":
			end = start + travel
			distance = d2cR
		case "L":
			end = start - travel
			distance = d2cL
		}
		cmdCounter := 0
		if travel >= distance {
			cmdCounter += 1
		}
		// if the distance is greater than 100 then we need to count how many extra 100s we crossed
		if travel > 100 {
			cmdCounter += (travel - distance) / 100
		}
		counter += cmdCounter
		start = end
	}
	return counter
}
