package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	println("Day 6")
	labMap, _, _, startPos := parseMap("../_input/day6.txt")
	findLoops(labMap, startPos, Pos{0, -1}, make(map[Pos][]Pos), make(map[Pos]bool))
}

func turnRight(dir Pos) Pos {
	return Pos{-dir.y, dir.x}
}

func findLoops(labMap map[Pos]string, currPos Pos, dir Pos, visited map[Pos][]Pos, loop_rock map[Pos]bool) {
	visited[currPos] = append(visited[currPos], dir)
	newPos := Pos{currPos.x + dir.x, currPos.y + dir.y}
	elem, inMap := labMap[newPos]
	if elem == "#" {
		dir = turnRight(dir)
		newPos = currPos
	} else if elem == "." {
		_, hasSeen := visited[newPos]
		if !hasSeen {
			labMap[newPos] = "#"
			if walkTillLoop(labMap, currPos, turnRight(dir), visited, make(map[Pos][]Pos)) {
				loop_rock[newPos] = true
			}
			labMap[newPos] = "."
		}
	} else if !inMap {
		fmt.Println(len(visited))   //P1
		fmt.Println(len(loop_rock)) //P2
		return
	}
	findLoops(labMap, newPos, dir, visited, loop_rock)
}

func walkTillLoop(labMap map[Pos]string, currPos Pos, dir Pos, visited map[Pos][]Pos, n_visited map[Pos][]Pos) bool {
	n_visited[currPos] = append(n_visited[currPos], dir)
	newPos := Pos{currPos.x + dir.x, currPos.y + dir.y}
	if slices.Contains(n_visited[newPos], dir) {
		return true
	}

	elem, inMap := labMap[newPos]
	if elem == "#" {
		dir = turnRight(dir)
		newPos = currPos
	} else if !inMap {
		return false
	}

	return walkTillLoop(labMap, newPos, dir, visited, n_visited)
}

func parseMap(file string) (map[Pos]string, int, int, Pos) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	labMap := make(map[Pos]string)
	maxY := len(lines)
	maxX := len(lines[0])
	startPos := Pos{0, 0}

	for i, line := range lines {
		for j, char := range line {
			labMap[Pos{j, i}] = string(char)
			if labMap[Pos{j, i}] != "." && labMap[Pos{j, i}] != "#" {
				startPos = Pos{j, i}
			}
		}
	}
	return labMap, maxX, maxY, startPos
}
