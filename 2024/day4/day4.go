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
	letters, maxX, maxY := readInput("../_input/day4.txt")
	findXMAS(letters, maxX, maxY)
	findXdashMAS(letters, maxX, maxY)

}

func readInput(file string) (map[Pos]string, int, int) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	letters := make(map[Pos]string)
	maxY := len(lines)
	maxX := len(lines[0])
	for i, line := range lines {
		for j, char := range line {
			letters[Pos{j, i}] = string(char)
		}
	}
	return letters, maxX, maxY
}

func findXdashMAS(letters map[Pos]string, maxX, maxY int) {
	count := 0
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			point := Pos{j, i}
			if letters[point] == "A" {
				x_letters := make([]string, 4)
				x_dirs := posNX()
				for i, dir := range x_dirs {
					n_p := Pos{point.x + dir.x, point.y + dir.y}
					x_letters[i] = letters[n_p]
				}
				if isXdashMAS(x_letters) {
					count += 1
				}
			}
		}
	}
	fmt.Println(count)
}

func isXdashMAS(letters []string) bool {
	m_count := 0
	s_count := 0
	for _, l := range letters {
		if l == "M" {
			m_count += 1
		} else if l == "S" {
			s_count += 1
		} else {
			return false
		}
	}
	if m_count != 2 {
		return false
	}
	if letters[0] == letters[2] {
		return false
	}
	return true
}

func findXMAS(letters map[Pos]string, maxX, maxY int) {
	count := 0
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			point := Pos{j, i}
			if letters[point] == "X" { // Could be the start of XMAS
				for _, dir := range posNDirections() {
					if walkWord(letters, point, dir, []string{"M", "A", "S"}) {
						count += 1
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func walkWord(letters map[Pos]string, point Pos, dir Pos, next_letters []string) bool {
	if len(next_letters) == 0 {
		return true
	}
	n_p := Pos{point.x + dir.x, point.y + dir.y}
	if letters[n_p] == next_letters[0] {
		return walkWord(letters, n_p, dir, next_letters[1:])
	}
	return false
}

func posNDirections() []Pos {
	return []Pos{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
}

func posNX() []Pos {
	return []Pos{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}
}
