package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ll, lr := readInput("../_input/day1.txt")
	sort.Ints(ll)
	sort.Ints(lr)

	part1(ll, lr)
	part2(ll, lr)
}

func part2(ll, lr []int) {
	sum := 0
	for _, val := range ll {
		freq := 0
		for _, trg := range lr {
			if trg == val {
				freq += 1
			}
		}
		sum += val * freq
	}
	fmt.Println(sum)
}

func part1(ll, lr []int) {
	distances := make([]int, len(ll))
	for i, _ := range distances {
		distances[i] = Abs(ll[i] - lr[i])
	}

	sum := 0
	for _, val := range distances {
		sum += val
	}

	fmt.Println(sum)

}

func first[T, U any](val T, _ U) T {
	return val
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func strsToInt(strs []string) []int {
	ints := make([]int, len(strs))
	for i, val := range strs {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		ints[i] = num
	}
	return ints
}

func readInput(file string) ([]int, []int) {
	contents, _ := os.ReadFile(file)
	str_pairs := strings.Split(string(contents), "\r\n")
	left_int_pairs := make([]int, len(str_pairs))
	right_int_pairs := make([]int, len(str_pairs))
	for i, val := range str_pairs {
		nums := strings.Split(val, "   ")
		ints := strsToInt(nums)
		left_int_pairs[i] = ints[0]
		right_int_pairs[i] = ints[1]
	}
	return left_int_pairs, right_int_pairs
}
