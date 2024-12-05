package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, print_lists := readInput("../_input/day5.txt")
	valid_prints := make([][]string, 0, len(print_lists))
	invalid_prints := make([][]string, 0, len(print_lists))
	for _, list := range print_lists {
		is_valid := true
		for j, el := range list {
			for _, after_num := range rules[el] {
				if slices.Contains(list[:j], after_num) {
					is_valid = false
				}
			}
		}
		if is_valid {
			valid_prints = append(valid_prints, list)
		} else {
			invalid_prints = append(invalid_prints, list)
		}
	}
	getMidSum(valid_prints)
	fixed_lists := make([][]string, len(invalid_prints))
	for i, list := range invalid_prints {
		fixed_lists[i] = reorderList(list, rules)
	}
	getMidSum(fixed_lists)
}

func reorderList(list []string, rules map[string][]string) []string {
	for i, el := range list {
		for _, after_num := range rules[el] {
			bad_num_found := elLocation(list[:i], after_num)
			if bad_num_found != -1 {
				list[bad_num_found] = list[i]
				list[i] = after_num
				return reorderList(list, rules)
			}
		}
	}
	return list
}

func elLocation(array []string, value string) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func getMidSum(prints [][]string) {
	acc := 0
	for _, vp := range prints {
		to_add, _ := strconv.Atoi(vp[len(vp)/2])
		acc += to_add
	}
	fmt.Println(acc)
}

func readInput(file string) (map[string][]string, [][]string) {
	contents, _ := os.ReadFile(file)
	input_parts := strings.Split(string(contents), "\n\n")
	page_order, print := input_parts[0], input_parts[1]
	page_orders := strings.Split(string(page_order), "\n")
	prints := strings.Split(string(print), "\n")

	rules := make(map[string][]string)
	print_lists := make([][]string, len(prints))
	for i, print := range prints {
		print_lists[i] = strings.Split(print, ",")
	}

	for _, order := range page_orders {
		bf := strings.Split(order, "|")
		rules[bf[0]] = append(rules[bf[0]], bf[1])
	}

	return rules, print_lists

}
