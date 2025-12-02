package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "ids.input"
	dat, _ := os.ReadFile(path)

	code := string(dat)
	items := strings.Split(code, ",")
	sum1, sum2 := 0, 0
	for _, item := range items {
		a, b := findInvalidCodeSums(parseRange(strings.TrimSpace(item)))
		sum1 += a
		sum2 += b
	}
	fmt.Println("Sum of invalid product codes:", sum1, sum2)
}

func parseRange(s string) (int, int) {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}

func findInvalidCodeSums(start, end int) (int, int) {
	sum1 := 0
	sum2 := 0
	for code := start; code <= end; code++ {
		sum1 += isValidCode1(code)
		sum2 += isValidCode2(code)
	}
	return sum1, sum2
}

func isValidCode1(code int) int {
	// code must be valid if it has odd length as you cant have a
	// sequence repeated twice
	codeStr := strconv.Itoa(code)
	codeLen := len(([]rune)(codeStr))
	if codeLen%2 != 0 {
		return 0
	}
	// check if first half matches second half
	halfLen := codeLen / 2
	return repeatedString(codeStr, halfLen)
}

func isValidCode2(code int) int {
	codeStr := strconv.Itoa(code)
	codeLen := len(([]rune)(codeStr))
	// maximimum length of string that can be repreated is half the length of the code
	maxRepeatLen := codeLen / 2
	// for each repeatable length, check if the string is made up of repeats
	for repeatLen := 1; repeatLen <= maxRepeatLen; repeatLen++ {
		// if there is a remainder, we cant have an exact repeat so skip
		if codeLen%repeatLen != 0 {
			continue
		}
		if repeatedString(codeStr, repeatLen) == 1 {
			return code
		}
	}
	return 0
}

func repeatedString(s string, repeatLen int) int {
	codeLen := len(([]rune)(s))
	firstItem := s[0:repeatLen]
	for i := 1; i < codeLen/repeatLen; i++ {
		nextItem := s[i*repeatLen : (i+1)*repeatLen]
		if nextItem != firstItem {
			return 0
		}
	}
	return 1
}
