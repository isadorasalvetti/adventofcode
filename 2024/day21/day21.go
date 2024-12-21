package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "../_sample/day21.txt"
	codes, code_num := parseMap(input)
	for i, code := range codes {
		current := 3
		seq := make([]rune, 0, 10)
		for _, char := range code {
			move_possibilities := moveKeypad(current, char)
			seq = append(seq, move_possibilities[0]...)
			seq = append(seq, 'A')
			current = char
		}
		res := moveDirpads(seq, 2)
		fmt.Println("Final bot", string(seq))
		fmt.Println(len(res), len(res)*code_num[i])
	}
}

// - Press same button as long as possible.
// - if in A, go up or right.
// - if changing, change to down.

//+---+---+---+
//|10 |11 |12 |
//+---+---+---+
//| 7 | 8 | 9 |
//+---+---+---+
//| 4 | 5 | 6 |
//+---+---+---+
//    | 2 | 3 |
//    +---+---+

//    +---+---+
//    | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+

func moveKeypad(pos int, target int) [][]rune {
	col := roundUp(target, 3) - roundUp(pos, 3)
	row := target - pos - (col * 3)

	return rowColSeq(row, col, pos, 1)
}

func solveDirpad(pos rune, seq []rune, res [][]rune) {
	target := seq[0]
	next_moves := moveDirpad(pos, target)
	for i, r := range res {
		for j, next_move := range next_moves {
			fmt.Println(i, j, r, next_move)
		}
	}
}

func moveDirpads(code []rune, how_many_bots int) []rune {
	current := 'A'
	seq := make([]rune, 0, 10)
	for _, char := range code {
		move_possibilities := moveDirpad(current, char)
		seq = append(seq, move_possibilities[0]...)
		seq = append(seq, 'A')
		current = char
	}
	fmt.Println("Bot", how_many_bots, string(seq))
	if how_many_bots == 0 {
		return seq
	} else {
		return moveDirpads(seq, how_many_bots-1)
	}
}

func moveDirpad(pos rune, target rune) [][]rune {
	var dir_pad = map[rune]int{
		'^': 1,
		'A': 2,
		'<': 3,
		'v': 4,
		'>': 5,
	}
	col := roundUp(dir_pad[target], 3) - roundUp(dir_pad[pos], 3)
	row := dir_pad[target] - dir_pad[pos] - (col * 3)
	return rowColSeq(row, col, dir_pad[pos], 0)
}

type Step struct {
	place int
	path  []rune
}

func rowColSeq(next []Step, target, dir_row, dir_col, blocked int) [][]rune {
	paths_found := make([][]rune, 0)
	arrows_row := []rune{'<', '>'}
	arrows_col := []rune{'v', '^'}

	for {
		if len(next) == 0 {
			return paths_found
		}
		head := next[0]
		tail := next[1:]

		if dir_row > 0 && roundUp(head.place+dir_row, 3) <= roundUp(target, 3) {
			next = append(next, Step{head.place + dir_row, append(head.path, arrows_row[dir_row])})
		} else if dir_row < 0 && roundUp(head.place+dir_row, 3) >= roundUp(target, 3) {
			next = append(next, Step{head.place + dir_row, append(head.path, arrows_row[dir_row+1])})
		}

		if dir_col > 0 && head.place+3 <= target {
			new_path := make([]rune, len(head.path)+1)
			copy(new_path, head.path)
			new_path[len(new_path)-1] = arrows_row[dir_row]
			next = append(next, Step{head.place + 3*dir_col + dir_row, new_path})
		} else if dir_col < 0 && head.place-3 >= target {
			new_path := make([]rune, len(head.path)+1)
			copy(new_path, head.path)
			new_path[len(new_path)-1] = arrows_row[dir_row+1]
			next = append(next, Step{head.place - 3*dir_row + dir_row, new_path})
		}

		if head.place+dir_row == target {
			head.path = append(head.path, arrows_row[dir_row+1])
		}

	}

	return [][]rune{}
}

func roundUp(num int, div int) int {
	return (num + div - 1) / div
}

func parseMap(file string) ([][]int, []int) {
	contents, _ := os.ReadFile(file)
	lines := strings.Split(string(contents), "\n")
	ls := make([][]int, len(lines))
	num_codes := make([]int, len(lines))
	for i, line := range lines {
		num_codes[i], _ = strconv.Atoi(line[:len(line)-1])
		ls[i] = make([]int, 4)
		for j, char := range line {
			if char == 'A' {
				ls[i][j] = 3
			} else if char == '0' {
				ls[i][j] = 2
			} else {
				ls[i][j] = int(char-'0') + 3
			}
		}
	}
	return ls, num_codes
}
