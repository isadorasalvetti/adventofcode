package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var moves_cache = map[toFrom][]rune{}
var bot_cache = map[string][]rune{}

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

func main() {
	acc := 0
	robots := 2
	input := "../_sample/day21.txt"

	codes, code_num := parseMap(input)
	for i, code := range codes {
		this_acc := 0
		current := 3
		for _, char := range code {
			snipp := append(moveKeypad(current, char), 'A')
			//fmt.Println("From numpad", string(snipp))
			snipp = solveDirpad('A', snipp, make([]rune, 0), robots)
			this_acc += len(snipp)
			current = char
			//fmt.Println("Done", i, char)
		}
		fmt.Println("Done line", code_num[i], this_acc)
		acc += this_acc * code_num[i]
	}
	fmt.Println(acc)
}

func solveDirpad(curr rune, seq []rune, res []rune, robots int) []rune {
	if robots == 0 {
		return seq
	}
	for {
		if len(seq) == 0 {
			return res
		}

		val, cached := bot_cache[string(seq)]
		if cached && false {
			return val
		}

		target := seq[0]
		next_move := append(moveDirpad(curr, target), 'A')
		res = append(res, solveDirpad('A', next_move, make([]rune, 0), robots-1)...)

		curr = target
		seq = seq[1:]
	}
}

func moveKeypad(pos int, target int) []rune {
	col := roundUp(target, 3) - roundUp(pos, 3)
	row := target - pos - (col * 3)

	return rowColSeq(row, col, pos, target, 1, false)
}

type toFrom struct {
	pos    rune
	target rune
}

func moveDirpad(pos rune, target rune) []rune {
	if pos == target {
		return make([]rune, 0)
	}

	val, cached := moves_cache[toFrom{pos, target}]
	if cached {
		return val
	}

	var dir_pad = map[rune]int{
		'^': 5,
		'A': 6,
		'<': 1,
		'v': 2,
		'>': 3,
	}
	col := roundUp(dir_pad[target], 3) - roundUp(dir_pad[pos], 3)
	row := dir_pad[target] - dir_pad[pos] - (col * 3)
	res_string := rowColSeq(row, col, dir_pad[pos], dir_pad[target], 4, true)
	moves_cache[toFrom{pos, target}] = res_string
	return rowColSeq(row, col, dir_pad[pos], dir_pad[target], 4, true)
}

func rowColSeq(row, col, pos, target, blocked int, rev bool) []rune {
	seq1 := make([]rune, 0, 5)
	seq2 := make([]rune, 0, 5)
	if col > 0 {
		for i := 0; i < col; i++ {
			seq1 = append(seq1, '^')
		}
	} else {
		for i := 0; i > col; i-- {
			seq1 = append(seq1, 'v')
		}
	}
	if row > 0 {
		for i := 0; i < row; i++ {
			seq2 = append(seq2, '>')
		}
	} else {
		for i := 0; i > row; i-- {
			seq2 = append(seq2, '<')
		}
	}

	if roundUp(pos, 3)-roundUp(blocked, 3) == 0 && blocked%3 == target%3 || roundUp(target, 3)-roundUp(blocked, 3) == 0 && blocked%3 == pos%3 {
		if target > pos || rev {
			return append(seq1, seq2...)
		} else {
			return append(seq2, seq1...)
		}
	}

	if len(seq1) == 0 || len(seq2) == 0 {
		return append(seq1, seq2...)
	}

	if row < 0 || col > 0 {
		return append(seq2, seq1...)
	}
	return append(seq1, seq2...)
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
