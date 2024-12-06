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
	labMap, _, _, startPos := parseMap("../_input/day6.txt")
	visited := make(map[Pos][]Pos)
	findLoops(labMap, startPos, Pos{0, -1}, visited, 0)
}

func turnRight(dir Pos) Pos {
	return Pos{-dir.y, dir.x}
}

func walk(labMap map[Pos]string, currPos Pos, dir Pos, visited map[Pos]Pos) map[Pos]Pos {
	visited[currPos] = dir
	newPos := Pos{currPos.x + dir.x, currPos.y + dir.y}
	elem, inMap := labMap[newPos]
	if elem == "#" {
		dir = turnRight(dir)
		newPos = currPos
	}
	if !inMap {
		return visited
	}
	return walk(labMap, newPos, dir, visited)
}

func findLoops(labMap map[Pos]string, currPos Pos, dir Pos, visited map[Pos][]Pos, loops int) {
	visited[currPos] = append(visited[currPos], dir)
	newPos := Pos{currPos.x + dir.x, currPos.y + dir.y}
	elem, inMap := labMap[newPos]
	if elem == "#" {
		dir = turnRight(dir)
		newPos = currPos
	} else {
		adir := turnRight(dir)
		if walkTillLoop(labMap, currPos, adir, visited) {
			loops += 1
		}
	}
	if !inMap {
		fmt.Println(len(visited)) //P1
		fmt.Println(loops)        //P2
		return
	}
	findLoops(labMap, newPos, dir, visited, loops)
}

func walkTillLoop(labMap map[Pos]string, currPos Pos, dir Pos, visited map[Pos][]Pos n_visited map[Pos][]Pos) bool {
	//visited[currPos] = append(visited[currPos], dir)
	newPos := Pos{currPos.x + dir.x, currPos.y + dir.y}
	elem, inMap := labMap[newPos]
	if elem == "#" {
		dir = turnRight(dir)
		newPos = currPos
	}
	if slices.Contains(visited[newPos], dir) {
		return true
	}
	if !inMap {
		return false
	}
	return walkTillLoop(labMap, newPos, dir, visited)
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
