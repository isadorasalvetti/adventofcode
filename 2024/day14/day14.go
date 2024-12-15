package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Robot struct {
	pos Pos
	vel Pos
}

type Pos struct {
	x int
	y int
}

type Quad struct {
	min Pos
	max Pos
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}

func mulPos(p1 Pos, mul int) Pos {
	return Pos{p1.x * mul, p1.y * mul}
}

func wrapPos(p1 Pos, size Pos) Pos {
	x := p1.x % size.x
	y := p1.y % size.y
	if x < 0 {
		x = size.x + x
	}

	if y < 0 {
		y = size.y + y
	}
	return Pos{x, y}
}

func main() {
	robots := parseInput("../_input/day14.txt")
	size := Pos{101, 103}
	iter := 100

	p1(robots, size, iter)
	p2(robots, size, iter+1)
}

func p2(robots []Robot, size Pos, offset int) {
	for v := offset; v < 20000; v++ {
		lines := make([][]int, size.y)
		for i, _ := range robots {
			robots[i].pos = wrapPos(addPos(robots[i].vel, robots[i].pos), size)
			lines[robots[i].pos.x] = append(lines[robots[i].pos.x], robots[i].pos.y)
		}
		for _, l := range lines {
			if len(l) > 8 {
				sort.Ints(l)
				for num := 1; num < len(l); num++ {
					if l[num]-l[num-1] != 1 {
						break
					}
					if num == len(l)-1 {
						//printBots(robots, size)
						fmt.Println(v)
						return
					}
				}
			}
		}
	}
}

func printBots(bots []Robot, size Pos) {
	bots_map := make(map[Pos]bool)
	for _, bot := range bots {
		bots_map[bot.pos] = true
	}
	for y := 0; y < size.y; y++ {
		for x := 0; x < size.x; x++ {
			if bots_map[Pos{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func p1(robots []Robot, size Pos, iter int) {
	for i, robot := range robots {
		robots[i].pos = wrapPos(addPos(mulPos(robot.vel, iter), robot.pos), size)
	}

	quadrants := []Quad{
		{Pos{0, 0}, Pos{size.x / 2, size.y / 2}},
		{Pos{size.x/2 + 1, 0}, Pos{size.x, size.y / 2}},
		{Pos{0, size.y/2 + 1}, Pos{size.x / 2, size.y}},
		{Pos{size.x/2 + 1, size.y/2 + 1}, Pos{size.x, size.y}},
	}

	quad_acc := make([]int, 4)

	for _, robot := range robots {
		for j, q := range quadrants {
			if robot.pos.x >= q.min.x && robot.pos.y >= q.min.y && robot.pos.x < q.max.x && robot.pos.y < q.max.y {
				quad_acc[j] += 1
				break
			}
		}
	}
	acc := 1
	for _, q := range quad_acc {
		acc *= q
	}
	fmt.Println(acc)
}

func parseInput(file string) []Robot {
	contents, _ := os.ReadFile(file)
	str_robots := strings.Split(string(contents), "\n")
	robots := make([]Robot, len(str_robots))
	for i, r := range str_robots {
		r_s := strings.Split(r[2:], " v=")
		p_s := strings.Split(r_s[0], ",")
		v_s := strings.Split(r_s[1], ",")

		px, _ := strconv.Atoi(p_s[0])
		py, _ := strconv.Atoi(p_s[1])
		vx, _ := strconv.Atoi(v_s[0])
		vy, _ := strconv.Atoi(v_s[1])

		robots[i] = Robot{Pos{px, py}, Pos{vx, vy}}
	}
	return robots
}
