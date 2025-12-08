package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord struct{ x, y, z int }
type jb struct {
	i int
	c coord
}

type pair struct{ a, b jb }

func (p pair) Distance() float64 {
	return math.Sqrt((math.Pow(float64(p.a.c.x-p.b.c.x), 2)) + (math.Pow(float64(p.a.c.y-p.b.c.y), 2)) + (math.Pow(float64(p.a.c.z-p.b.c.z), 2)))
}

func main() {
	dat, _ := os.ReadFile("jbs.input")
	rows := strings.Split(string(dat), "\n")
	junctionBoxes := []jb{}
	for i, row := range rows {
		vals := strings.Split(row, ",")
		jb := jb{i: i, c: coord{x: parseInt(vals[0]), y: parseInt(vals[1]), z: parseInt(vals[2])}}
		junctionBoxes = append(junctionBoxes, jb)
	}

	pairMap := map[pair]bool{}
	for i, jb := range junctionBoxes {
		for j, otherJb := range junctionBoxes {
			if i == j {
				continue
			}
			if _, ok := pairMap[pair{a: otherJb, b: jb}]; ok {
				continue
			}
			pairMap[pair{a: jb, b: otherJb}] = true
		}
	}
	pairs := []pair{}
	for p := range pairMap {
		pairs = append(pairs, p)
	}
	// sort by distance
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Distance() < pairs[j].Distance()
	})
	part1Circuits := []map[jb]bool{}
	N := 1000
	for _, pwd := range pairs[:N] {
		part1Circuits = append(part1Circuits, map[jb]bool{pwd.a: true, pwd.b: true})
	}
	combinedCircuits := combineCircuits(part1Circuits)
	c1len := len(combinedCircuits[0])
	c2len := len(combinedCircuits[1])
	c3len := len(combinedCircuits[2])

	a, b := combineCircuits2(pairs, len(junctionBoxes))
	fmt.Println("part 1:", c1len*c2len*c3len)
	fmt.Println("part 2:", a, b, a.c.x*b.c.x)

}

func combineCircuits(circuits []map[jb]bool) []map[jb]bool {
	newCircuits := []map[jb]bool{circuits[0]}
	for _, circuit := range circuits[1:] {
		overlap := false
		foundCircuit := 0
		for j, newCircuit := range newCircuits {
			for jb := range circuit {
				if _, ok := newCircuit[jb]; ok {
					overlap = true
					foundCircuit = j
					break
				}
			}

		}
		if overlap {
			for jb := range circuit {
				newCircuits[foundCircuit][jb] = true
			}
		} else {
			newCircuits = append(newCircuits, circuit)
		}
	}
	if len(newCircuits) == len(circuits) {
		sort.Slice(newCircuits, func(i, j int) bool {
			return len(newCircuits[i]) > len(newCircuits[j])
		})
		return newCircuits
	} else {
		return combineCircuits(newCircuits)
	}
}

func combineCircuits2(pairs []pair, totalJbs int) (jb, jb) {
	newCircuits := map[jb]map[jb]bool{pairs[0].a: {pairs[0].a: true, pairs[0].b: true}}
	for _, pair := range pairs[1:] {
		// find if a is in a circuit
		var cAKey, cBKey *jb
		for j, newCircuit := range newCircuits {
			if _, ok := newCircuit[pair.a]; ok {
				cAKey = &j
			}
			if _, ok := newCircuit[pair.b]; ok {
				cBKey = &j
			}
		}
		if cAKey != nil || cBKey != nil {
			if cAKey != nil && cBKey != nil && cAKey != cBKey {
				// merge circuits
				for jb := range newCircuits[*cBKey] {
					newCircuits[*cAKey][jb] = true
				}
				delete(newCircuits, *cBKey)
				l := len(newCircuits[*cAKey])
				if l == totalJbs {
					return pair.a, pair.b
				}
			} else if cAKey != nil {
				newCircuits[*cAKey][pair.b] = true
				l := len(newCircuits[*cAKey])
				if l == totalJbs {
					return pair.a, pair.b
				}
			} else if cBKey != nil {
				newCircuits[*cBKey][pair.a] = true
				l := len(newCircuits[*cBKey])
				if l == totalJbs {
					return pair.a, pair.b
				}
			}

		} else {
			newCircuits[pair.a] = map[jb]bool{pair.a: true, pair.b: true}
		}

	}
	return jb{}, jb{}
}

func parseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
