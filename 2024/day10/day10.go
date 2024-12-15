package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	nodes, zeros, _, _ := parseMap("../_input/day10.txt")
	acc1 := 0
	acc2 := 0
	for _, zero := range zeros {
		s1, s2 := run_trail(nodes, []Pos{zero}, make(map[Pos]bool), 0)
		acc1 += s1
		acc2 += s2
	}
	fmt.Println(acc1)
	fmt.Println(acc2)
}

func run_trail(mymap map[Pos]int, to_visit []Pos, ends map[Pos]bool, score int) (int, int) {
	if len(to_visit) == 0 {
		return len(ends), score
	}
	head := to_visit[0]
	to_visit = to_visit[1:]
	for _, dir := range dirs() {
		new_location := addPos(dir, head)
		if mymap[new_location]-mymap[head] == 1 {
			if mymap[new_location] == 9 {
				ends[new_location] = true
				score += 1
			} else {
				to_visit = append(to_visit, new_location)
			}
		}
	}
	return run_trail(mymap, to_visit, ends, score)
}

func dirs() []Pos {
	return []Pos{Pos{0, 1}, Pos{0, -1}, Pos{1, 0}, Pos{-1, 0}}
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}

func parseMap(file string) (map[Pos]int, []Pos, int, int) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	height_map := make(map[Pos]int)
	zeros := make([]Pos, 0)
	maxY := len(lines)
	maxX := len(lines[0])

	for i, line := range lines {
		for j, char := range line {
			node_height := int(char - '0')
			height_map[Pos{j, i}] = node_height
			if node_height == 0 {
				zeros = append(zeros, Pos{j, i})
			}
		}
	}
	return height_map, zeros, maxX, maxY
}
