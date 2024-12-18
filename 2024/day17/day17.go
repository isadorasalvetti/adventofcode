package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	registers, instructions := parseInput("../_input/day17.txt")
	out := runInstructions(registers, instructions)
	fmt.Print("P1: ", out[0])
	for _, o := range out[1:] {
		fmt.Print(",", o)
	}
	fmt.Println()

	guess := brute_force(instructions)
	fmt.Println("P2: ", guess)

}

func brute_force(instructions []int) int {
	correct_digits := 0
	curr := 10000000000000
	iter := 0
	correct_digits = 0
	for {
		interval := intpow(8, (14 - correct_digits))
		curr += interval
		registers := []int{curr, 0, 0}
		out := runInstructions(registers, instructions)
		new_correct_digs := 0
		if len(out) == len(instructions) {
			for i := len(out) - 1; i >= 0; i-- {
				if out[i] == instructions[i] {
					new_correct_digs += 1
				} else {
					break
				}
			}
		}
		correct_digits = new_correct_digs
		if correct_digits == len(instructions) {
			return curr
		}
		if curr > 100000000000000000 {
			break
		}
		iter += 1
	}
	return 0
}

func runInstructions(registers, instructions []int) []int {
	out := make([]int, 0)

	for i := 0; i < len(instructions); i += 2 {
		inst := instructions[i]
		lit := instructions[i+1]

		res := -1
		//fmt.Println("Lit:", lit, "Reg:", registers)
		switch inst {
		case 0:
			//fmt.Println("Made Register A: RegA/2^comb")
			res = registers[0] / intpow(2, getComb(lit, registers))
			registers[0] = res
		case 1:
			//fmt.Println("Made Register B: RegB^lit")
			res = registers[1] ^ lit
			registers[1] = res
		case 2:
			//fmt.Println("Made Register B: RegB%8")
			res = getComb(lit, registers) % 8
			registers[1] = res
		case 3:
			//fmt.Println("Change i")
			if registers[0] != 0 {
				i = lit - 2
			}
		case 4:
			//fmt.Println("Made Register B: RegB^RegC")
			res = registers[1] ^ registers[2]
			registers[1] = res
		case 5:
			//fmt.Println("Out %8")
			out = append(out, getComb(lit, registers)%8)
		case 6:
			//fmt.Println("Made Register B: RegA/2^comb")
			res = registers[0] / intpow(2, getComb(lit, registers))
			registers[1] = res
		case 7:
			//fmt.Println("Made Register C: RegA/2^comb")
			res = registers[0] / intpow(2, getComb(lit, registers))
			registers[2] = res
		}
	}
	return out
}

func getComb(lit int, registers []int) int {
	switch lit {
	case 0, 1, 2, 3:
		return lit
	case 4:
		//fmt.Println("Using Reg A")
		return registers[0]
	case 5:
		//fmt.Println("Using Reg B")
		return registers[1]
	case 6:
		//fmt.Println("Using Reg C")
		return registers[2]
	}
	panic("Bad combo")
}

func intpow(num, pow int) int {
	res := 1
	for pow > 0 {
		res *= num
		pow -= 1
	}
	return res
}

func parseInput(file string) ([]int, []int) {
	contents, _ := os.ReadFile(file)
	parts := strings.Split(string(contents), "\n\n")
	register_lines := strings.Split(parts[0], "\n")
	registers := make([]int, len(register_lines))

	instructions_lines := strings.Split(parts[1][9:], ",")
	instructions := make([]int, len(instructions_lines))

	for i, r_line := range register_lines {
		r_line_split := strings.Split(r_line, ": ")
		registers[i], _ = strconv.Atoi(r_line_split[1])
	}

	for i, instr := range instructions_lines {
		instructions[i], _ = strconv.Atoi(instr)
	}
	return registers, instructions
}
