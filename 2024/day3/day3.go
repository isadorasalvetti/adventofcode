package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	num_pairs1 := readInput("../_input/day3.txt")
	sum_mults(num_pairs1)

	num_pairs2 := readInput2("../_input/day3.txt")
	sum_mults(num_pairs2)
}

func sum_mults(nums [][]int) {
	acc := 0
	for _, val := range nums {
		acc += val[0] * val[1]
	}
	fmt.Println(acc)
}

func readInput(file string) [][]int {
	contents, _ := os.ReadFile(file)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(string(contents), -1)
	num_pairs := make([][]int, len(matches))
	for i, val := range matches {
		num_pairs[i] = getNums(val)
	}
	return num_pairs
}

func readInput2(file string) [][]int {
	contents, _ := os.ReadFile(file)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	matches := re.FindAllString(string(contents), -1)
	num_pairs := make([][]int, 0, len(matches))
	count_next := true
	for _, val := range matches {
		if val == "do()" {
			count_next = true
		} else if val == "don't()" {
			count_next = false
		} else if count_next {
			num_pairs = append(num_pairs, getNums(val))
		}
	}
	return num_pairs
}

func getNums(num_string string) []int {
	str_nums := strings.Split(num_string[4:len(num_string)-1], ",")
	str_ints := make([]int, 2)
	for j, num := range str_nums {
		str_ints[j], _ = strconv.Atoi(num)
	}
	return str_ints
}
