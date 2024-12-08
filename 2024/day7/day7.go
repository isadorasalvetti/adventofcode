package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Eq struct {
	res  int
	nums []int
}

func main() {
	println("Day 7")
	lines := parseLines("../_input/day7.txt")
	acc := 0
	for _, line := range lines {
		if doOp(line.nums[0], line.nums[1:], line.res) {
			acc += line.res
		}
	}
	fmt.Println(acc)
	acc = 0
	for _, line := range lines {
		if doOp2(line.nums[0], line.nums[1:], line.res) {
			acc += line.res
		}
	}
	fmt.Println(acc)
}

func doOp(lastres int, nums []int, res int) bool {
	if len(nums) == 0 {
		return lastres == res
	} else {
		return doOp(lastres+nums[0], nums[1:], res) || doOp(lastres*nums[0], nums[1:], res)
	}
}

func doOp2(lastres int, nums []int, res int) bool {
	if len(nums) == 0 {
		return lastres == res
	} else {
		len_s := lenLoop(nums[0])
		return doOp2(lastres+nums[0], nums[1:], res) || doOp2(lastres*nums[0], nums[1:], res) || doOp2((lastres*PowInts(10, len_s))+nums[0], nums[1:], res)
	}
}

func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func parseLines(file string) []Eq {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	line_list := make([]Eq, len(lines))

	for i, line := range lines {
		els := strings.Split(line, ": ")
		res, _ := strconv.Atoi(els[0])
		nums := strings.Split(els[1], " ")
		nums_list := make([]int, len(nums))
		for j, num := range nums {
			nums_list[j], _ = strconv.Atoi(num)
		}
		line_list[i] = Eq{res, nums_list}
	}
	return line_list
}
