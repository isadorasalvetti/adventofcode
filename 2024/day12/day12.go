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

type Fence struct {
	s1 Pos
	s2 Pos
}

func main() {
	garden, _, _ := parseInput("../_input/day12.txt")
	areas(garden)
}

func areas(garden map[Pos]rune) map[int][]Pos {
	areas := make(map[int][]Pos)

	acc1 := 0
	acc2 := 0
	for pos, val := range garden { //deleted entries will not be iterated
		plot, perimeter := getPlot(garden, make(map[Pos]bool), val, pos, make(map[Fence]int))
		total_sides := getContEdges(perimeter)
		acc1 += len(plot) * len(perimeter)
		acc2 += len(plot) * total_sides
	}
	fmt.Println(acc1)
	fmt.Println(acc2)
	return areas
}

func getContEdges(edges map[Fence]int) int {
	next_edge_id_counter := 1
	for f, id := range edges {
		if id == 0 {
			edges[f] = next_edge_id_counter
			next_edge_id_counter += 1
		}

		ed_base := contEdgesDir(f)
		for _, ed := range []Pos{ed_base, negPos(ed_base)} {
			ne1 := Fence{addPos(f.s1, ed), addPos(f.s2, ed)}
			for {
				_, exists := edges[ne1]
				if exists {
					edges[ne1] = edges[f]
					ne1 = Fence{addPos(ne1.s1, ed), addPos(ne1.s2, ed)}
				} else {
					break
				}
			}
		}
	}
	return next_edge_id_counter - 1
}

func contEdgesDir(edge Fence) Pos {
	edge_perp_dir := subPos(edge.s1, edge.s2)
	pd := turn90(edge_perp_dir) // Edge parallel direction
	return pd
}

func getPlot(garden map[Pos]rune, plot map[Pos]bool, plot_type rune, pos Pos, perimeter map[Fence]int) (map[Pos]bool, map[Fence]int) {
	delete(garden, pos)

	plot[pos] = true
	dirs := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		to_look := addPos(pos, d)
		p, in_map := garden[addPos(pos, d)]
		if in_map && p == plot_type {
			plot[to_look] = true
			plot, perimeter = getPlot(garden, plot, plot_type, to_look, perimeter)
		} else {
			_, in_plot := plot[to_look]
			if !in_plot {
				perimeter[Fence{pos, to_look}] = 0
			}
		}
	}
	return plot, perimeter
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}

func subPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x - p2.x, p2.y - p1.y}
}

func turn90(dir Pos) Pos {
	return Pos{-dir.y, dir.x}
}

func negPos(dir Pos) Pos {
	return Pos{-dir.x, -dir.y}
}

func parseInput(file string) (map[Pos]rune, int, int) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	garden := make(map[Pos]rune)
	maxY := len(lines)
	maxX := len(lines[0])

	for i, line := range lines {
		for j, char := range line {
			garden[Pos{i, j}] = char
		}
	}
	return garden, maxX, maxY
}
