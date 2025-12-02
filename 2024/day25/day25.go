package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "../_input/day25.txt"
	locks, keys := parseInput(input)
	fit := 0
	for _, lock := range locks {
		for _, key := range keys {
			key_fits := true
			for i := range lock {
				if key[i] < lock[i] {
					//fmt.Println(key[i], lock[i])
					key_fits = false
					break
				}
			}
			//fmt.Println(lock, key, key_fits)
			if key_fits {
				fit += 1
			}
		}
	}
	fmt.Println(fit)

}

func parseInput(file string) ([][]int, [][]int) {
	contents, _ := os.ReadFile(file)
	locks_s := strings.Split(string(contents), "\n\n")
	locks := [][]int{}
	keys := [][]int{}
	for _, lock := range locks_s {
		lock_lines := strings.Split(lock, "\n")
		if lock_lines[0] == "#####" { // Key
			locks = append(locks, make([]int, 5))
			for j := range lock_lines[0] {
				for h := len(lock_lines) - 1; h >= 0; h-- {
					if lock_lines[h][j] == '#' {
						locks[len(locks)-1][j] = h + 1
						break
					}
				}
			}
		} else {
			keys = append(keys, make([]int, 5))
			for j := range lock_lines[0] {
				for h := 0; h < len(lock_lines); h++ {
					if lock_lines[h][j] == '#' {
						keys[len(keys)-1][j] = h
						break
					}
				}
			}
		}
	}
	return locks, keys
}
