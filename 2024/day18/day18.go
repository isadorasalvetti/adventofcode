package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Step struct {
	pos  Pos
	time int
}

func main() {
	falling_bits, size, time := parseInput("../_sample/day18.txt")
	bit_falling_rock := make(map[Pos]int)
	for i, rock := range falling_bits {
		bit_falling_rock[rock] = i
	}
	steps := walkMem(Step{Pos{0, 0}, 0}, make(map[Pos]bool), bit_falling_rock, time, size)
	fmt.Println(steps)
}

func walkMem(curr Step, visited map[Pos]bool, blocked map[Pos]int, time int, size int) int {
	visited[curr.pos] = true

	if curr.pos.x == size && curr.pos.y == size {
		return curr.time
	}

	for _, d := range dirs() {
		poss_next := addPos(curr.pos, d)
		when, is_blocked := blocked[poss_next]
		if visited[poss_next] || (is_blocked && when < time) || !inBound(poss_next, size) {
			continue
		}
		walkMem(Step{poss_next, curr.time + 1}, visited, blocked, time, size)
	}
	return -1
}

func inBound(pos Pos, size int) bool {
	return pos.x <= size && pos.x >= 0 && pos.y <= size && pos.y >= 0
}

func dirs() []Pos {
	return []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
}

func addPos(a, b Pos) Pos {
	return Pos{a.x + b.x, a.y + b.y}
}

func parseInput(file string) ([]Pos, int, int) {
	contents, _ := os.ReadFile(file)
	parts := strings.Split(string(contents), "\n\n")
	lines := strings.Split(string(parts[0]), "\n")
	nums := strings.Split(string(parts[1]), ",")

	size, _ := strconv.Atoi(nums[0])
	time, _ := strconv.Atoi(nums[1])
	bits := make([]Pos, len(lines))

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		bits[i] = Pos{x, y}
	}
	return bits, size, time
}
