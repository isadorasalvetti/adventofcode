package main

import (
	"container/heap"
	"fmt"
	"math"
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
	falling_bits, size, time := parseInput("../_input/day18.txt")
	bit_falling_rock := make(map[Pos]int)
	for i, rock := range falling_bits {
		bit_falling_rock[rock] = i
	}
	steps := walkMem([]Path{{[]Pos{{0, 0}}, 0}}, make(map[Pos]bool), bit_falling_rock, time, size, 0)
	fmt.Println("P1", steps.time)
	blocked_time := 0
	for {
		if steps.time == -1 {
			fmt.Println("P2", blocked_time, falling_bits[blocked_time])
			break
		}
		blocked_time = pathIsBlocked(steps, bit_falling_rock, blocked_time)
		steps = walkMem([]Path{{[]Pos{{0, 0}}, 0}}, make(map[Pos]bool), bit_falling_rock, blocked_time, size, 0)
	}
}

func walkMem(next PathHeap, visited map[Pos]bool, blocked map[Pos]int, time, size, recur int) Path {
	for {
		if next.Len() == 0 {
			return Path{make([]Pos, 0), -1}
		}
		curr := heap.Pop(&next).(Path)
		curr_pos := curr.path[len(curr.path)-1]
		if visited[curr_pos] {
			continue
		}
		visited[curr_pos] = true

		if curr_pos.x == size && curr_pos.y == size {
			return curr
		}

		for _, d := range dirs() {
			poss_next := addPos(curr_pos, d)
			when, is_blocked := blocked[poss_next]
			if visited[poss_next] || (is_blocked && when <= time) || !inBound(poss_next, size) {
				continue
			}
			new_array := make([]Pos, len(curr.path)+1)
			copy(new_array, curr.path)
			new_array[len(curr.path)] = poss_next
			next.Push(Path{new_array, curr.time + 1})
		}
		recur += 1
	}
}

func pathIsBlocked(path Path, blocked map[Pos]int, time int) int {
	soonest_blocked := math.MaxInt64
	for _, spot := range path.path {
		if blocked[spot] > time {
			soonest_blocked = min(soonest_blocked, blocked[spot])
		}
	}
	return soonest_blocked
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

type PathHeap []Path

func (h PathHeap) Len() int {
	return len(h)
}

func (h PathHeap) Less(i int, j int) bool {
	return h[i].time < h[j].time
}

func (h PathHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (s *PathHeap) Push(x any) {
	h := *s
	h = append(h, x.(Path))
	*s = h
}

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
