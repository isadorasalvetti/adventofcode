package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	_, towels, orders := parseInput("../_input/day19.txt")
	fmt.Println(towels)
	will_split := make(map[string]bool)
	acc := 0
	for _, o := range orders {
		s := fitsTowel(towels, will_split, o)
		fmt.Println("DONE:", o, s)
		if s {
			acc += 1
		}
	}
	fmt.Println(acc)
}

func fitsTowel(towels []string, will_fit map[string]bool, order string) bool {
	in, ws := will_fit[order]
	if in {
		return ws
	}
	for _, t := range towels {
		if len(t) <= len(order) && order[:len(t)] == t {
			if len(order) == len(t) {
				return true
			}
			rest_fits := fitsTowel(towels, will_fit, order[len(t):])
			will_fit[order[len(t):]] = rest_fits
			if rest_fits {
				return true
			}
		}
	}
	return false
}

func trySplit(towels map[string]bool, will_split map[string]bool, order string, split int) bool {
	in, ws := will_split[order]
	if in {
		return ws
	}

	head := order[:split]
	tail := order[split:]
	if len(head) == 0 {
		panic("head should never be empty")
	}
	//fmt.Println("Got", head, tail, split, len(order))

	if towels[head] {
		if len(tail) == 0 {
			return true
		}

		tail_split := trySplit(towels, will_split, tail, 1)
		will_split[tail] = tail_split
		if tail_split {
			//fmt.Println("Tail split worked")
			return true
		}
		//fmt.Println("Bad head/tail split, continue", split, order, len(order))
	}
	if len(tail) == 0 {
		return false
	}
	return trySplit(towels, will_split, order, split+1)
}

func parseInput(file string) (map[string]bool, []string, []string) {
	contents, _ := os.ReadFile(file)
	parts := strings.Split(string(contents), "\n\n")
	towels := strings.Split(string(parts[0]), ", ")
	orders := strings.Split(string(parts[1]), "\n")

	towels_map := make(map[string]bool)
	for _, t := range towels {
		towels_map[t] = true
	}

	return towels_map, towels, orders
}
