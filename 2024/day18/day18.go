package main

import (
	"container/heap"
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

type Path struct {
	path []Pos
	time int
}

func main() {
	falling_bits, size, time := parseInput("../_sample/day18.txt")
	bit_falling_rock := make(map[Pos]int)
	for i, rock := range falling_bits {
		bit_falling_rock[rock] = i
	}

	steps := walkMem([]Step{{Pos{0, 0}, 0}}, make(map[Pos]bool), bit_falling_rock, time, size, 0)
	max_time := exitPathExistsTime([]Step{{Pos{0, 0}, 0}}, make(map[Pos]bool), bit_falling_rock, size, 0)

	fmt.Println(steps)
	fmt.Println(max_time)

}

func walkMem(next StepHeap, visited map[Pos]bool, blocked map[Pos]int, time, size, recur int) int {
	for {
		curr := heap.Pop(&next).(Step)
		if visited[curr.pos] {
			continue
		}

		visited[curr.pos] = true

		if curr.pos.x == size && curr.pos.y == size {
			return curr.time
		}

		for _, d := range dirs() {
			poss_next := addPos(curr.pos, d)
			when, is_blocked := blocked[poss_next]
			if visited[poss_next] || (is_blocked && when < time) || !inBound(poss_next, size) {
				continue
			} else {
				next.Push(Step{poss_next, curr.time + 1})
			}
		}
		recur += 1
	}
}

func exitPathExistsTime(next StepHeap, visited map[Pos]bool, blocked map[Pos]int, size, recur int) Step {
	rock_blocked := Step{Pos{0, 0}, 0}
	for {
		if len(next) == 0 {
			return rock_blocked
		}
		curr := heap.Pop(&next).(Step)
		when_blocked := blocked[curr.pos]
		if when_blocked > 0 && rock_blocked.time < when_blocked {
			rock_blocked = Step{curr.pos, when_blocked}
			continue
		}

		if visited[curr.pos] {
			continue
		}
		visited[curr.pos] = true

		for _, d := range dirs() {
			poss_next := addPos(curr.pos, d)
			if visited[poss_next] || !inBound(poss_next, size) {
				continue
			}
			next.Push(Step{poss_next, curr.time + 1})
		}
		recur += 1
	}
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

type StepHeap []Step

func (h StepHeap) Len() int {
	return len(h)
}

func (h StepHeap) Less(i int, j int) bool {
	return h[i].time < h[j].time
}

func (h StepHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (s *StepHeap) Push(x any) {
	h := *s
	h = append(h, x.(Step))
	*s = h
}

func (h *StepHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
