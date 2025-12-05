package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ingredientCodeRanges []ingredientCodeRange

func (r ingredientCodeRanges) Count() int {
	totals := 0
	for _, icr := range r {
		totals += icr.end - icr.start + 1
	}
	return totals
}

type ingredientCodeRange struct {
	start, end int
}

func (icr ingredientCodeRange) union(rng ingredientCodeRange) []ingredientCodeRange {
	// if ranges are contiguous or overlapping, return a single range
	if rng.start <= icr.end && rng.end >= icr.start {
		newRng := ingredientCodeRange{start: min(icr.start, rng.start), end: max(icr.end, rng.end)}
		return []ingredientCodeRange{newRng}
	}
	// else return both ranges
	return []ingredientCodeRange{icr, rng}
}

func main() {
	freshItemsStrs := loadAndSplitFile("fresh.input", "\n")
	freshItems := make([]ingredientCodeRange, len(freshItemsStrs))
	for i, item := range freshItemsStrs {
		start, end := parseRange(strings.TrimSpace(item))
		freshItems[i] = ingredientCodeRange{start: start, end: end}
	}
	newRngs := unlapRanges(freshItems)
	fmt.Println("total number of fresh items:", newRngs.Count())
}

func unlapRanges(ranges ingredientCodeRanges) ingredientCodeRanges {
	if len(ranges) <= 1 {
		return ranges
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	new := ingredientCodeRanges{}
	for i := 0; i < len(ranges); {
		fir := ranges[i]
		if i < len(ranges)-1 {
			nexr := ranges[i+1]
			mergedRanges := fir.union(nexr)
			if len(mergedRanges) == 1 {
				new = append(new, mergedRanges...)
				i += 2
			} else {
				new = append(new, fir)
				i++
			}
		} else {
			new = append(new, fir)
			break
		}
	}
	if new.Count() == ranges.Count() {
		return new
	} else {
		return unlapRanges(new)
	}
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
