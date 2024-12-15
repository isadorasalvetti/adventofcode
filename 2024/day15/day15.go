package main

import (
	"fmt"
	"os"
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
	for _, dir := range dirs {
		to_print := false
		var to_update map[Pos]rune
		if (nodes[addPos(bot_pos, dirToPos(dir))] == '[' || nodes[addPos(bot_pos, dirToPos(dir))] == ']') && (dir == '>' || dir == '<') {
			to_print = true
			fmt.Println(string(dir))
			printBots(nodes, bot_pos, Pos{x_s, y_s})
		}
		bot_pos, to_update, _ = move2(nodes, bot_pos, dir)
		for pos, val := range to_update {
			nodes[pos] = val
		}
		if to_print {
			printBots(nodes, bot_pos, Pos{x_s, y_s})
		}
	}
	getScore(nodes)
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

func move2(nodes map[Pos]rune, curr_pos Pos, dir rune) (Pos, map[Pos]rune, bool) {
	to_update := make(map[Pos]rune)

	new_pos := addPos(curr_pos, dirToPos(dir))
	if nodes[new_pos] == '.' {
		return new_pos, to_update, true
	} else if nodes[new_pos] == '#' {
		return curr_pos, make(map[Pos]rune), false
	}

	if dir == '<' || dir == '>' {
		new_pos_ad := addPos(new_pos, dirToPos(dir))
		target, _, moved := move2(nodes, new_pos_ad, dir)
		if moved {
			nodes[target] = nodes[new_pos_ad]
			nodes[new_pos_ad] = nodes[new_pos]
			nodes[new_pos] = '.'
			return new_pos, to_update, true
		}
	} else {
		var new_pos_ad Pos
		if nodes[new_pos] == '[' {
			new_pos_ad = addPos(new_pos, Pos{1, 0})
		} else {
			new_pos_ad = addPos(new_pos, Pos{-1, 0})
		}
		target, extra_update, moved := move2(nodes, new_pos, dir)
		for pos, val := range extra_update {
			to_update[pos] = val
		}
		target_ad, extra_update, moved_ad := move2(nodes, new_pos_ad, dir)
		for pos, val := range extra_update {
			to_update[pos] = val
		}

		if moved && moved_ad {
			to_update[target] = nodes[new_pos]
			to_update[target_ad] = nodes[new_pos_ad]
			to_update[new_pos] = '.'
			to_update[new_pos_ad] = '.'
			return new_pos, to_update, true
		}
	}
	return curr_pos, make(map[Pos]rune), false
}

func printBots(nodes map[Pos]rune, bot Pos, size Pos) {
	for y := bot.y - 6; y < bot.y+6; y++ {
		for x := bot.x - 6; x < bot.x+6; x++ {
			if bot.x == x && bot.y == y {
				fmt.Print("@")
				continue
			}
			fmt.Print(string(nodes[Pos{x, y}]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func dirToPos(dir_char rune) Pos {
	return map[rune]Pos{'<': {-1, 0}, 'v': {0, 1}, '>': {1, 0}, '^': {0, -1}}[dir_char]
}

func addPos(p1 Pos, p2 Pos) Pos {
	return Pos{p1.x + p2.x, p2.y + p1.y}
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
