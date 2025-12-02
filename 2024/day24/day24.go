package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := "../_input/day24.txt"
	operations, vals, _, _, zs := parseInput(input)
	slices.Sort(zs)
	slices.Reverse(zs)

	findPairs(operations, zs)
	resolveOps(operations, vals)

	final := []rune{}
	for _, z := range zs {
		if vals[z] {
			final = append(final, '1')
		} else {
			final = append(final, '0')
		}
	}
	fmt.Println(string(final))
	fmt.Println(strconv.ParseInt(string(final), 2, 64))

}

func getPath(operations map[string][]string, next []string) [][]string {
	path := [][]string{}
	for {
		if len(next) == 0 {
			break
		}
		curr := next[0]
		op := operations[curr]
		path = append(path, append(op, curr))
		next = next[1:]
		if strings.HasPrefix(op[0], "x") || strings.HasPrefix(op[0], "y") {
			continue
		}
		if strings.HasPrefix(op[2], "y") || strings.HasPrefix(op[1], "x") {
			continue
		}
		next = append(next, op[0], op[2])
	}
	return path
}

func findPairs(operations map[string][]string, zs []string) {
	bad := map[string][]string{}
	for rs, op := range operations {
		if op[1] == "XOR" {
			if op[0][0] == 'x' || op[0][0] == 'y' {
				if rs != "z00" {
					found := false
					for _, op_s := range operations {
						if (op_s[0] == rs || op_s[2] == rs) && op_s[1] == "XOR" {
							found = true
							break
						}
					}
					if !found {
						bad[rs] = op
					}
				}
			} else {
				if rs[0] != 'z' { // valid: tgd XOR rvg -> z01
					bad[rs] = op
				}
			}
		} else {
			if rs[0] == 'z' && (rs[1] != '4' && rs[2] != '5') { // valid: tgd XOR rvg -> z01
				bad[rs] = op
			} else if op[1] == "AND" && op[2] != "x00" {
				found := false
				for _, op_s := range operations {
					if (op_s[0] == rs || op_s[2] == rs) && op_s[1] == "OR" {
						found = true
						break
					}
				}
				if !found {
					bad[rs] = op
				}
			}
		}
	}
	fmt.Println(bad)
	bad_wires := slices.Collect(maps.Keys(bad))
	slices.Sort(bad_wires)
	println(strings.Join(bad_wires, ","))
}

func resolveOps(operations map[string][]string, vals map[string]bool) {
	for len(operations) > 0 {
		for res, op := range operations {
			res1, in_map1 := vals[op[0]]
			res2, in_map2 := vals[op[2]]
			operator := op[1]

			if in_map1 && in_map2 {
				if operator == "AND" {
					vals[res] = res1 && res2
				} else if operator == "OR" {
					vals[res] = res1 || res2
				} else if operator == "XOR" {
					vals[res] = !(res1 == res2)
				}
				delete(operations, res)
				//fmt.Println("Set", res, vals[res], op, res1, res2)
			}
		}
	}
}

func parseInput(file string) (map[string][]string, map[string]bool, []string, []string, []string) {
	contents, _ := os.ReadFile(file)
	parts := strings.Split(string(contents), "\n\n")
	vals := strings.Split(string(parts[0]), "\n")
	ops_str := strings.Split(string(parts[1]), "\n")
	zs := []string{}
	xs := []string{}
	ys := []string{}

	operations := make(map[string][]string)
	completed := make(map[string]bool)
	for _, line := range vals {
		p := strings.Split(line, ": ")
		conv, _ := strconv.Atoi(p[1])
		completed[p[0]] = conv != 0
		if strings.HasPrefix(p[1], "x") {
			xs = append(xs, p[1])
		} else if strings.HasPrefix(p[1], "y") {
			ys = append(ys, p[1])
		}
	}

	for _, line := range ops_str {
		p := strings.Split(line, " -> ")
		operations[p[1]] = strings.Split(p[0], " ")
		if strings.HasPrefix(p[1], "z") {
			zs = append(zs, p[1])
		}
	}

	return operations, completed, xs, ys, zs
}
