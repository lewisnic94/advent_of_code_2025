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
	N := 1000

	newCircuits := map[jb]map[jb]bool{pairs[0].a: {pairs[0].a: true, pairs[0].b: true}}
	totalJbs := len(junctionBoxes)
	for j, pair := range pairs[1:] {
		// find if a or b is in a circuit
		var cAKey, cBKey *jb
		for j, newCircuit := range newCircuits {
			if _, ok := newCircuit[pair.a]; ok {
				cAKey = &j
			}
			if _, ok := newCircuit[pair.b]; ok {
				cBKey = &j
			}
		}
		var l int = 2
		if cAKey != nil && cBKey != nil && cAKey != cBKey {
			// merge circuits
			for jb := range newCircuits[*cBKey] {
				newCircuits[*cAKey][jb] = true
			}
			delete(newCircuits, *cBKey)
			l = len(newCircuits[*cAKey])
		} else if cAKey != nil {
			newCircuits[*cAKey][pair.b] = true
			l = len(newCircuits[*cAKey])
		} else if cBKey != nil {
			newCircuits[*cBKey][pair.a] = true
			l = len(newCircuits[*cBKey])
		} else {
			newCircuits[pair.a] = map[jb]bool{pair.a: true, pair.b: true}
		}
		if l == totalJbs {
			fmt.Println("part 2:", pair.a.c.x*pair.b.c.x)
			return
		} else if j == N-1 {
			// get the 3 largest circuits
			circuits := []map[jb]bool{}
			for _, c := range newCircuits {
				circuits = append(circuits, c)
			}
			sort.Slice(circuits, func(i, j int) bool {
				return len(circuits[i]) > len(circuits[j])
			})
			fmt.Println("part 1:", len(circuits[0])*len(circuits[1])*len(circuits[2]))
		}
	}
}

func parseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
