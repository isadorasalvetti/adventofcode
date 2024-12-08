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
	antenas, maxX, maxY := parseMap("../_input/day8.txt")
	antinodes := make(map[Pos]bool)
	for _, a := range antenas {
		getAntinodes(a, maxX, maxY, antinodes)
	}
	fmt.Println(len(antinodes))

	for _, a := range antenas {
		getInlineAntinodes(a, maxX, maxY, antinodes)
	}
	fmt.Println(len(antinodes))

}

func getAntinodes(antenas []Pos, maxX int, maxY int, antinodes map[Pos]bool) {
	for _, antena := range antenas {
		for _, antena2 := range antenas {
			if antena2 == antena {
				continue
			}
			dist := distPos(antena, antena2)
			a1, a2 := addPos(antena, dist), addPos(antena2, minusPos(dist))
			if inMap(a1, maxX, maxY) {
				antinodes[a1] = true
			}
			if inMap(a2, maxX, maxY) {
				antinodes[a2] = true
			}
		}
	}
}

func getInlineAntinodes(antenas []Pos, maxX int, maxY int, antinodes map[Pos]bool) {
	for _, antena := range antenas {
		for _, antena2 := range antenas {
			if antena2 == antena {
				continue
			}
			interval := notNormalize(distPos(antena, antena2))
			a1 := antena
			for inMap(a1, maxX, maxY) {
				antinodes[a1] = true
				a1 = addPos(a1, interval)
			}

			a1 = antena
			for inMap(a1, maxX, maxY) {
				antinodes[a1] = true
				a1 = addPos(a1, minusPos(interval))
			}
		}
	}
}

func inMap(p1 Pos, maxX, maxY int) bool {
	return p1.x < maxX && p1.x >= 0 && p1.y < maxY && p1.y >= 0
}

func distPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x - p2.x, p1.y - p2.y}
}

func notNormalize(dir Pos) Pos {
	for i := dir.x; i > 0; i-- {
		if dir.x%i == 0 && dir.y%i == 0 {
			return Pos{dir.x / i, dir.y / i}
		}
	}
	return dir
}

func minusPos(p1 Pos) Pos {
	return Pos{-p1.x, -p1.y}
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}

func parseMap(file string) (map[string][]Pos, int, int) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	labMap := make(map[string][]Pos)
	maxY := len(lines)
	maxX := len(lines[0])

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "." || string(char) == "#" {
				continue
			}
			labMap[string(char)] = append(labMap[string(char)], Pos{j, i})
		}
	}
	return labMap, maxX, maxY
}
