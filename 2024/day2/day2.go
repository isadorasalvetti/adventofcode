package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rows := readInput("../_input/day2.txt")
	part1 := 0
	part2 := 0
	for _, val := range rows {
		if isSafe(val, true) || isSafe(val, false) {
			part1 += 1
		}
		r_val := reverseInts(val)
		if isSafeTolerance(val, true) || isSafeTolerance(val, false) || isSafeTolerance(r_val, true) || isSafeTolerance(r_val, false) {
			part2 += 1
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func isSafe(row []int, updown bool) bool {
	last_dig := row[0]
	for _, val := range row[1:] {
		dif := val - last_dig
		if updown && (dif > 3 || dif < 1) {
			return false
		}
		if !updown && (dif < -3 || dif > -1) {
			return false
		}
		last_dig = val
	}
	return true
}

func isSafeTolerance(row []int, updown bool) bool {
	last_dig := row[0]
	tolerance := 1
	for _, val := range row[1:] {
		discard := false
		dif := val - last_dig
		if updown && (dif > 3 || dif < 1) {
			if tolerance > 0 {
				tolerance -= 1
				discard = true
			} else {
				return false
			}
		}
		if !updown && (dif < -3 || dif > -1) {
			if tolerance > 0 {
				tolerance -= 1
				discard = true
			} else {
				return false
			}
		}
		if !discard {
			last_dig = val
		}
	}
	return true
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

func readInput(file string) [][]int {
	contents, _ := os.ReadFile(file)
	str_lines := strings.Split(string(contents), "\r\n")
	num_rows := make([][]int, len(str_lines))
	for i, line := range str_lines {
		nums := strings.Split(line, " ")
		ints := strsToInt(nums)
		num_rows[i] = ints
	}
	return num_rows
}
