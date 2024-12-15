package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := parseInput("../_input/day11.txt")
	rounds := 0
	for rounds < 75 {
		next_round := make(map[int]int)
		for stone, occ := range stones {
			if stone == 0 {
				next_round[1] += occ
			} else {
				stone_str := strconv.Itoa(stone)
				if len(stone_str)%2 == 0 {
					s1, _ := strconv.Atoi(stone_str[:len(stone_str)/2])
					s2, _ := strconv.Atoi(stone_str[len(stone_str)/2:])
					next_round[s1] += occ
					next_round[s2] += occ
				} else {
					next_round[stone*2024] += occ
				}
			}
		}
		stones = next_round
		rounds += 1

		if rounds == 25 {
			countStones(stones)
		}

		if rounds == 75 {
			countStones(stones)
		}
	}
}

func countStones(stones map[int]int) {
	acc := 0
	for _, val := range stones {
		acc += val
	}

	fmt.Println(acc)
}

func parseInput(file string) map[int]int {
	contents, _ := os.ReadFile(file)
	nums := strings.Split(string(contents), " ")
	nodes := make(map[int]int)
	for _, n := range nums {
		num_int, _ := strconv.Atoi(n)
		nodes[num_int] += 1
	}
	return nodes
}
