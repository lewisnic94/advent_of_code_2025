package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "rolls.input"
	dat, _ := os.ReadFile(path)
	code := string(dat)
	items := strings.Split(code, "\n")
	rollsMatrix := make([][]int, len(items))
	for i, item := range items {
		item := strings.Replace(item, ".", "0", -1)
		item = strings.Replace(item, "@", "1", -1)
		nums := ConvertToNumList(item)
		rollsMatrix[i] = nums
	}
	movableCount := 0

	for {
		var count int
		rollsMatrix, count = performPass(rollsMatrix)
		if count == 0 {
			break
		}
		movableCount += count
	}
	fmt.Println("Number of movable positions:", movableCount)
}

func performPass(matrix [][]int) ([][]int, int) {
	movableCount := 0
	for j := 0; j < len(matrix); j++ { // y
		for i := 0; i < len(matrix[j]); i++ { // x
			if matrix[j][i] == 0 {
				continue
			}
			count := countSurroundingRolls(matrix, j, i)
			if count < 4 {
				movableCount++
				matrix[j][i] = 0
			}
		}
	}
	return matrix, movableCount
}

func countSurroundingRolls(matrix [][]int, y, x int) int {

	moves := [8][2]int{ // X and Y offsets
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	count := 0
	for _, move := range moves {

		Ynew := y + move[0]
		Xnew := x + move[1]
		if Ynew < 0 || Ynew == len(matrix) {
			continue
		}
		if Xnew < 0 || Xnew == len(matrix[Ynew]) {
			continue
		}
		count += matrix[Ynew][Xnew]
	}

	return count
}

func ConvertToNumList(s string) []int {
	numsStrs := strings.SplitN(s, "", -1)
	nums := make([]int, len(numsStrs))
	for i, str := range numsStrs {
		n, _ := strconv.Atoi(str)
		nums[i] = n
	}
	return nums
}
