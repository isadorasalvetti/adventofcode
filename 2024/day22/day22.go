package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seq struct {
	a int
	b int
	c int
	d int
}

func main() {
	input := "../_input/day22.txt"
	seeds := parseMap(input)
	steps := 2000
	acc := 0
	all_seqs := make(map[Seq][]int)
	for _, seed := range seeds {
		last_nums := make([]int, steps+1)
		last_nums[0] = seed % 10
		seqs := make(map[Seq]int)
		for i := 1; i < steps+1; i++ {
			seed = doStep(seed)
			last_nums[i] = seed % 10
		}
		for i := 4; i < len(last_nums); i++ {
			a := last_nums[i-3] - last_nums[i-4]
			b := last_nums[i-2] - last_nums[i-3]
			c := last_nums[i-1] - last_nums[i-2]
			d := last_nums[i] - last_nums[i-1]
			_, seen := seqs[Seq{a, b, c, d}]
			if !seen {
				seqs[Seq{a, b, c, d}] = last_nums[i]
			}
		}
		for key, val := range seqs {
			all_seqs[key] = append(all_seqs[key], val)
		}
		acc += seed
	}
	fmt.Println(acc)
	fmt.Println(getBestSeq(all_seqs))
}

func getBestSeq(all_seqs map[Seq][]int) (Seq, int) {
	var best Seq
	var nums []int
	best_val := 0
	for key, val := range all_seqs {
		acc := 0
		for _, x := range val {
			acc += x
		}
		if best_val < acc {
			best_val = acc
			best = key
			nums = val
		}
	}
	fmt.Println(nums)
	return best, best_val
}

func doStep(secret int) int {
	res := secret * 64
	secret = res ^ secret
	secret = secret % 16777216

	res = secret / 32
	secret = res ^ secret
	secret = secret % 16777216

	res = secret * 2048
	secret = res ^ secret
	secret = secret % 16777216

	return secret
}

func parseMap(file string) []int {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	num_codes := make([]int, len(lines))
	for i, line := range lines {
		num_codes[i], _ = strconv.Atoi(line)
	}
	return num_codes
}
