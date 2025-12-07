package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("beams.input")
	rows := strings.Split(string(dat), "\n")
	incomingBeams := map[int]int{strings.Index(rows[0], "S"): 1}
	splitCount, totalTimelines := 0, 0
	for _, row := range rows {
		outgoingBeams := map[int]int{}
		for idx, count := range incomingBeams {
			if []rune(row)[idx] == '^' {
				outgoingBeams[idx-1] += count
				outgoingBeams[idx+1] += count
				splitCount++
			} else {
				outgoingBeams[idx] += count
			}
		}
		incomingBeams = outgoingBeams
	}
	for _, count := range incomingBeams {
		totalTimelines += count
	}
	fmt.Println("number of timelines:", totalTimelines, "number of splits:", splitCount)
}
