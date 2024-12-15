package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	a_x     int
	a_y     int
	b_x     int
	b_y     int
	prize_x int
	prize_y int
}

func main() {
	machines := parseInput("../_input/day13.txt")
	conv := 10000000000000

	acc1 := 0
	acc2 := 0
	for _, machine := range machines {
		acc1 += solve(machine, 0)
		acc2 += solve(machine, conv)
	}
	fmt.Println(acc1, "P1")
	fmt.Println(acc2, "P2")
}

func solve(machine Machine, conv int) int {
	div := (machine.b_x * machine.a_y) - (machine.b_y * machine.a_x)
	if div == 0 {
		return 0
	}
	num := ((machine.prize_x + conv) * machine.a_y) - ((machine.prize_y + conv) * machine.a_x)

	if num%div != 0 {
		return 0
	}
	pb := num / div

	if ((machine.prize_x+conv)-pb*machine.b_x)%machine.a_x != 0 {
		return 0
	}
	pa := ((machine.prize_x + conv) - pb*machine.b_x) / machine.a_x

	if pa < 0 && pb < 0 {
		return 0
	}

	if conv == 0 && pa > 100 && pb > 100 {
		return 0
	}
	if machine.prize_x+conv != pb*machine.b_x+pa*machine.a_x {
		fmt.Println("MATH AINT MATHING")
	}
	return pa*3 + pb
}

func parseInput(file string) []Machine {
	contents, _ := os.ReadFile(file)
	str_machines := strings.Split(string(contents), "\n\n")
	machines := make([]Machine, len(str_machines))
	for i, m := range str_machines {
		m_lines := strings.Split(m, "\n")
		a_line := strings.Split(strings.Split(m_lines[0], "Button A: X+")[1], ", Y+")
		a_x, _ := strconv.Atoi(a_line[0])
		a_y, _ := strconv.Atoi(a_line[1])

		b_line := strings.Split(strings.Split(m_lines[1], "Button B: X+")[1], ", Y+")
		b_x, _ := strconv.Atoi(b_line[0])
		b_y, _ := strconv.Atoi(b_line[1])

		m_line := strings.Split(strings.Split(m_lines[2], "Prize: X=")[1], ", Y=")
		m_x, _ := strconv.Atoi(m_line[0])
		m_y, _ := strconv.Atoi(m_line[1])
		machines[i] = Machine{a_x, a_y, b_x, b_y, m_x, m_y}
	}
	return machines
}
