package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	input := "../_input/day15.txt"
	nodes, bot_pos, dirs, x_s, x_y := parseMap1(input)
	processBoxes1(nodes, bot_pos, dirs, x_s, x_y)

	nodes, bot_pos, dirs, x_s, x_y = parseMap2(input)
	processBoxes2(nodes, bot_pos, dirs, x_s, x_y)
}

func processBoxes1(nodes map[Pos]rune, bot_pos Pos, dirs string, x_s int, y_s int) {
	for _, dir := range dirs {
		bot_pos, _ = move(nodes, bot_pos, dir)
		//printBots(nodes, bot_pos, Pos{x_s, y_s})
	}
	getScore(nodes)
}

func processBoxes2(nodes map[Pos]rune, bot_pos Pos, dirs string, x_s int, y_s int) {
	for i, dir := range dirs {
		var to_update map[Pos]rune
		var bot_pos_add Pos
		bot_pos, bot_pos_add, to_update, _ = move2(nodes, bot_pos, dir)

		if false && (dir == 'v' || dir == '^') && len(to_update) > 0 {
			printBots(nodes, bot_pos, Pos{x_s, y_s})
		}

		update_y_move(nodes, to_update, dirToPos(dir))
		nodes[bot_pos] = '.'
		nodes[bot_pos_add] = '.'

		if false && (dir == 'v' || dir == '^') && len(to_update) > 0 {
			fmt.Println(string(dir), " ", to_update)
			printBots(nodes, bot_pos, Pos{x_s, y_s})
		}

		if false && i == 1249 {
			getScore(nodes)
			fmt.Println(string(dir))
			printBots(nodes, bot_pos, Pos{x_s, y_s})
		}
	}
	//printBots(nodes, bot_pos, Pos{x_s, y_s})
	getScore(nodes)
}

func update_y_move(nodes map[Pos]rune, to_update map[Pos]rune, dir Pos) {
	to_update_sorted := make([]Pos, 0)
	for pos, _ := range to_update {
		to_update_sorted = append(to_update_sorted, pos)
	}

	sort.Slice(to_update_sorted, func(i, j int) bool {
		if dir.y == -1 {
			return to_update_sorted[i].y < to_update_sorted[j].y
		} else {
			return to_update_sorted[i].y > to_update_sorted[j].y
		}
	})
	for _, pos := range to_update_sorted {
		nodes[addPos(pos, dir)] = nodes[pos]
		nodes[pos] = '.'
	}
}

func getScore(nodes map[Pos]rune) {
	acc := 0
	for pos, node := range nodes {
		if node == 'O' || node == '[' {
			acc += gpsCoord(pos)
		}
	}
	fmt.Println(acc)
}

func gpsCoord(pos Pos) int {
	return pos.x + pos.y*100
}

func move(nodes map[Pos]rune, curr_pos Pos, dir rune) (Pos, bool) {
	new_pos := addPos(curr_pos, dirToPos(dir))
	if nodes[new_pos] == '.' {
		return new_pos, true
	} else if nodes[new_pos] == 'O' {
		target, will_move := move(nodes, new_pos, dir)
		if will_move {
			nodes[new_pos] = '.'
			nodes[target] = 'O'
			return new_pos, true
		} else {
			return curr_pos, false
		}
	}
	return curr_pos, false
}

func move2(nodes map[Pos]rune, curr_pos Pos, dir rune) (Pos, Pos, map[Pos]rune, bool) {
	to_update := make(map[Pos]rune)

	new_pos := addPos(curr_pos, dirToPos(dir))
	if nodes[new_pos] == '.' {
		return new_pos, new_pos, to_update, true
	}
	if nodes[new_pos] == '[' || nodes[new_pos] == ']' {
		if dir == '<' || dir == '>' {
			new_pos_ad := addPos(new_pos, dirToPos(dir))
			target, _, _, moved := move2(nodes, new_pos_ad, dir)
			if moved {
				nodes[target] = nodes[new_pos_ad]
				nodes[new_pos_ad] = nodes[new_pos]
				nodes[new_pos] = '.'
				return new_pos, new_pos, to_update, true
			}
		} else {
			var new_pos_ad Pos
			if nodes[new_pos] == '[' {
				new_pos_ad = addPos(new_pos, Pos{1, 0})
			} else {
				new_pos_ad = addPos(new_pos, Pos{-1, 0})
			}
			_, _, extra_update, moved := move2(nodes, new_pos, dir)
			_, _, extra_update_ad, moved_ad := move2(nodes, new_pos_ad, dir)
			for pos, val := range extra_update {
				to_update[pos] = val
			}
			for pos, val := range extra_update_ad {
				to_update[pos] = val
			}
			if moved && moved_ad {
				//to_update[target] = nodes[new_pos]
				//to_update[target_ad] = nodes[new_pos_ad]
				to_update[new_pos] = '.'
				to_update[new_pos_ad] = '.'
				return new_pos, new_pos_ad, to_update, true
			}
		}
	}
	return curr_pos, curr_pos, make(map[Pos]rune), false
}

func inMap(map_ map[Pos]rune, key Pos) bool {
	_, in_map := map_[key]
	return in_map
}

func printBots(nodes map[Pos]rune, bot Pos, size Pos) {
	for y := 0; y < size.y; y++ {
		for x := 0; x < size.x; x++ {
			if bot.x == x && bot.y == y {
				fmt.Print("@")
				continue
			}
			fmt.Print(string(nodes[Pos{x, y}]))
		}
		fmt.Println()
	}
}

func dirToPos(dir_char rune) Pos {
	return map[rune]Pos{'<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}, '^': {0, -1}}[dir_char]
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
}

func minusPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x - p2.x, p1.y - p2.y}
}

func parseMap1(file string) (map[Pos]rune, Pos, string, int, int) {
	contents, _ := os.ReadFile(file)
	parts := strings.Split(string(contents), "\n\n")
	lines := strings.Split(string(parts[0]), "\n")
	warehouse_map := make(map[Pos]rune)
	maxY := len(lines)
	maxX := len(lines[0])
	var robot Pos

	for i, line := range lines {
		for j, char := range line {
			if char == '@' {
				robot = Pos{j, i}
				char = '.'
			}
			warehouse_map[Pos{j, i}] = char
		}
	}
	return warehouse_map, robot, parts[1], maxX, maxY
}

func parseMap2(file string) (map[Pos]rune, Pos, string, int, int) {
	contents, _ := os.ReadFile(file)
	parts := strings.Split(string(contents), "\n\n")
	lines := strings.Split(string(parts[0]), "\n")
	warehouse_map := make(map[Pos]rune)
	maxY := len(lines)
	maxX := len(lines[0])
	var robot Pos

	for i, line := range lines {
		for j, char := range line {
			if char == 'O' {
				warehouse_map[Pos{j * 2, i}] = '['
				warehouse_map[Pos{j*2 + 1, i}] = ']'
			} else {
				if char == '@' {
					robot = Pos{j * 2, i}
					char = '.'
				}
				warehouse_map[Pos{j * 2, i}] = char
				warehouse_map[Pos{j*2 + 1, i}] = char
			}
		}
	}
	return warehouse_map, robot, parts[1], maxX * 2, maxY
}
