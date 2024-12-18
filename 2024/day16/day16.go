package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Step struct {
	pos Pos
	dir Pos
}

type ScorePath struct {
	head  Step
	path  map[Step]bool
	score int
}

func main() {
	input := "../_input/day16.txt"
	maze, start, end, _, _ := parseMap(input)
	start_dir := Pos{1, 0}
	next := &StepHeap{ScoreStep{Step{start, start_dir}, 0}}
	heap.Init(next)

	best_score := travelMaze(next, maze, make(map[Step]bool), end)
	fmt.Println(best_score)

	next_path1 := &PathHeap{ScorePath{Step{start, start_dir}, make(map[Step]bool), 0}}
	next_path2 := &PathHeap{ScorePath{Step{end, turn90(start_dir)}, make(map[Step]bool), 0},
		ScorePath{Step{end, turn90(turn90(start_dir))}, make(map[Step]bool), 0}}
	heap.Init(next_path1)
	heap.Init(next_path2)

	dist_to_start := inBestPath(next_path1, maze, make(map[Step]int), best_score)
	dist_to_end := inBestPath(next_path2, maze, make(map[Step]int), best_score)

	bestPos := make(map[Pos]bool)
	for step, dist := range dist_to_start {
		if dist+dist_to_end[Step{step.pos, turn90(turn90(step.dir))}] == best_score {
			bestPos[step.pos] = true
		}
	}
	fmt.Println(len(bestPos))
}

func travelMaze(next *StepHeap, maze map[Pos]rune, visited map[Step]bool, end Pos) int {
	if next.Len() == 0 {
		fmt.Println("No end.")
		return 0
	}

	curr := heap.Pop(next).(ScoreStep)
	if curr.step.pos == end {
		return curr.score
	}
	visited[curr.step] = true

	for next_step, next_score := range nextOptions(maze, curr.step, curr.score) {
		in_visited := visited[next_step]
		if !in_visited {
			heap.Push(next, ScoreStep{next_step, next_score})
		}
	}

	return travelMaze(next, maze, visited, end)
}

func inBestPath(next *PathHeap, maze map[Pos]rune, visited_dist map[Step]int, best_score int) map[Step]int {
	for next.Len() > 0 {
		curr := heap.Pop(next).(ScorePath)
		prev_score, seen := visited_dist[curr.head]
		if !seen || prev_score > curr.score {
			visited_dist[curr.head] = curr.score
		} else {
			continue
		}
		for next_step, next_score := range nextOptions(maze, curr.head, curr.score) {
			if next_score <= best_score && !curr.path[next_step] {
				heap.Push(next, ScorePath{next_step, clone(curr.path), next_score})
			}
		}
	}
	return visited_dist
}

func clone(to_clone map[Step]bool) map[Step]bool {
	new := make(map[Step]bool)
	for k, v := range to_clone {
		new[k] = v
	}
	return new
}

func nextOptions(maze map[Pos]rune, step Step, score int) map[Step]int {
	poss_steps := make(map[Step]int)
	poss_steps[Step{addPos(step.pos, step.dir), step.dir}] = score + 1
	poss_steps[Step{step.pos, turn90(step.dir)}] = score + 1000
	poss_steps[Step{step.pos, turn90(turn90(turn90(step.dir)))}] = score + 1000

	for step, _ := range poss_steps {
		if maze[step.pos] != '.' {
			delete(poss_steps, step)
		}
	}
	return poss_steps
}

func parseMap(file string) (map[Pos]rune, Pos, Pos, int, int) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	maze := make(map[Pos]rune)
	maxY := len(lines)
	maxX := len(lines[0])
	var start Pos
	var end Pos

	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				start = Pos{j, i}
				char = '.'
			}
			if char == 'E' {
				end = Pos{j, i}
				char = '.'
			}
			maze[Pos{j, i}] = char
		}
	}
	return maze, start, end, maxX, maxY
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}

func turn90(dir Pos) Pos {
	return Pos{-dir.y, dir.x}
}

type PathHeap []ScorePath

func (h PathHeap) Len() int {
	return len(h)
}

func (h PathHeap) Less(i int, j int) bool {
	return h[i].score < h[j].score
}

func (h PathHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (s *PathHeap) Push(x any) {
	h := *s
	h = append(h, x.(ScorePath))
	*s = h
}

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type StepHeap []ScoreStep

type ScoreStep struct {
	step  Step
	score int
}

func (h StepHeap) Len() int {
	return len(h)
}

func (h StepHeap) Less(i int, j int) bool {
	return h[i].score < h[j].score
}

func (h StepHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (s *StepHeap) Push(x any) {
	h := *s
	h = append(h, x.(ScoreStep))
	*s = h
}

func (h *StepHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
