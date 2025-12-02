package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type Conn struct {
	pc1 string
	pc2 string
	pc3 string
}

func main() {
	input := "../_input/day23.txt"
	graph := parseInput(input)
	p1(graph)
	p2(graph)
}

type clickStep struct {
	click string
	next  string
}

func p2(graph map[string][]string) {
	clicks := map[string]bool{}
	seen := map[string]bool{}
	next := []clickStep{}
	for pc1, pcns := range graph {
		for _, pc2 := range pcns {
			click_l := []string{pc1, pc2}
			slices.Sort(click_l)
			click_s := clickStep{click_l[0], click_l[1]}
			next = append(next, click_s)
		}
	}

	for len(next) > 0 {
		next_step := next[0]
		next = next[1:]
		if seen[next_step.click] {
			continue
		}

		seen[next_step.click] = true
		click_pc_l := strings.Split(next_step.click, ",")
		click_pc_l = append(click_pc_l, next_step.next)
		slices.Sort(click_pc_l)
		full_click := strings.Join(click_pc_l, ",")

		for _, poss_pc := range graph[next_step.next] {
			if slices.Contains(click_pc_l, poss_pc) {
				continue
			}
			in_click := true
			for _, click_pc := range click_pc_l {
				if !slices.Contains(graph[poss_pc], click_pc) {
					in_click = false
				}
			}

			if in_click {
				next = append(next, clickStep{full_click, poss_pc})
			} else {
				clicks[full_click] = true
			}
		}
	}
	max_click := ""
	for click, _ := range clicks {
		if len(click) > len(max_click) {
			max_click = click
		}
	}
	fmt.Println(max_click)
}

func Pop(dict map[clickStep]bool) (clickStep, bool) {
	for key, value := range dict {
		delete(dict, key)
		return key, value
	}
	panic("Poop on empty dict")
}

func p1(graph map[string][]string) {
	conn := map[Conn]bool{}
	for pc1, pc2_list := range graph {
		if !strings.HasPrefix(pc1, "t") {
			continue
		}
		for _, pc2 := range pc2_list {
			for _, pc3 := range graph[pc2] {
				if pc3 != pc1 && slices.Contains(graph[pc3], pc1) {
					pcs := []string{pc1, pc2, pc3}
					sort.Strings(pcs)
					conn[Conn{pcs[0], pcs[1], pcs[2]}] = true
				}
			}
		}
	}
	fmt.Println(len(conn))

}

func parseInput(file string) map[string][]string {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	graph := make(map[string][]string)
	for _, line := range lines {
		pcs := strings.Split(line, "-")
		graph[pcs[0]] = append(graph[pcs[0]], pcs[1])
		graph[pcs[1]] = append(graph[pcs[1]], pcs[0])
	}
	return graph
}
