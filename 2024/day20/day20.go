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
	input := "../_input/day20.txt"
	maze, start, end := parseMap(input)
	dists := make(map[Pos]int)
	scoreMaze(maze, dists, start, end, 0)

	//fmt.Println(countCheat(maze, dists, start, 3))

	acc1 := 0
	acc2 := 0
	savings1 := make(map[int]int)
	savings2 := make(map[int]int)
	for pos, _ := range dists {
		for _, saved := range countCheat(maze, dists, pos, 2) {
			savings1[saved] += 1
			if saved >= 100 {
				acc1 += 1
			}
		}
		for _, saved := range countCheat(maze, dists, pos, 20) {
			if saved >= 50 {
				savings2[saved] += 1
			}
			if saved >= 100 {
				acc2 += 1
			}
		}
	}

	//fmt.Println(savings1)
	//fmt.Println(savings2)
	fmt.Println(acc1)
	fmt.Println(acc2)
}

func countCheat(maze map[Pos]rune, dists map[Pos]int, pos Pos, how_many int) map[Pos]int {
	to_return := make(map[Pos]int)
	poss := make(map[Pos]int)
	allposs := make(map[Pos]int)
	for _, dir := range dirs() { // Cheat starts on first positions around pos
		poss[addPos(dir, pos)] = 1
		allposs[addPos(dir, pos)] = 1
	}
	for len(poss) > 0 {
		n_pos, steps := Pop(poss)
		spot, exists := maze[n_pos]
		if spot == '.' { // Exits wall, cheat ends
			time_save := dists[n_pos] - dists[pos] - steps
			//fmt.Println("Exit: ", n_pos, steps, time_save)
			if time_save > to_return[n_pos] {
				to_return[n_pos] = time_save
			}
		}
		if exists && steps < how_many {
			//fmt.Println("Look again: ", n_pos, steps)
			for _, dir := range dirs() {
				nn_poss := addPos(dir, n_pos)
				when, seen := allposs[nn_poss]
				if !seen || when > steps {
					poss[nn_poss] = steps + 1
					allposs[nn_poss] = steps + 1
				}
			}
		}
	}
	return to_return
}

func Pop(dict map[Pos]int) (Pos, int) {
	for key, value := range dict {
		delete(dict, key)
		return key, value
	}
	panic("Poop on empty dict")
}

func scoreMaze(maze map[Pos]rune, dists map[Pos]int, next Pos, end Pos, time int) {
	dists[next] = time
	if next == end {
		return
	}
	for _, dir := range dirs() {
		poss_next := addPos(dir, next)
		_, visited := dists[poss_next]
		if maze[poss_next] == '.' && !visited {
			scoreMaze(maze, dists, poss_next, end, time+1)
			return
		}
	}
}

func dirs() []Pos {
	return []Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
}

func parseMap(file string) (map[Pos]rune, Pos, Pos) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	maze := make(map[Pos]rune)
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
	return maze, start, end
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}
