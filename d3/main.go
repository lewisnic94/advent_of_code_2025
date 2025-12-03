package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "joltage.input"
	dat, _ := os.ReadFile(path)
	code := string(dat)
	items := strings.Split(code, "\n")
	joltageSum := 0
	for _, item := range items {
		numsStrs := strings.SplitN(item, "", 100)
		ints := make([]int, len(numsStrs))
		for i, s := range numsStrs {
			n, _ := strconv.Atoi(s)
			ints[i] = n
		}
		joltage := maximiseJoltage(ints, 12)
		joltageSum += joltage
	}
	fmt.Println("Sum of max joltage:", joltageSum)
}

func maximiseJoltage(nums []int, n int) int {
	cursor, joltage := 0, 0
	for i := range n {
		max, offset := 0, 0
		numsSubset := nums[cursor : len(nums)-n+i+1]
		for j, num := range numsSubset {
			if num > max {
				max, offset = numsSubset[j], j
			}
		}
		joltage += nums[cursor+offset] * int(math.Pow10(n-(i+1)))
		cursor += offset + 1
	}
	return joltage
}
