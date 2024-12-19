package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	_, towels, orders := parseInput("../_input/day19.txt")
	will_split := make(map[string]int)
	acc1 := 0
	acc2 := 0
	for _, o := range orders {
		s := fitsTowel(towels, will_split, o, 0)
		fmt.Println("DONE:", o, s)
		if s > 0 {
			acc1 += 1
		}
		acc2 += s
	}
	fmt.Println(acc1)
	fmt.Println(acc2)
}

func fitsTowel(towels []string, will_fit map[string]int, order string, ret int) int {
	ws, in := will_fit[order]
	if in {
		return ws
	}

	for _, t := range towels {
		rest := strings.TrimSuffix(order, t)
		if rest == order {
			continue
		}
		if len(rest) == 0 {
			ret += 1
			continue
		}
		rest_fits := fitsTowel(towels, will_fit, rest, 0)
		//fmt.Println(t, rest, rest_fits)
		will_fit[rest] = rest_fits
		ret += rest_fits
	}
	return ret
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
