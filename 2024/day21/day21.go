package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	acc := 0
	input := "../_input/day21.txt"
	codes, code_num := parseMap(input)
	for i, code := range codes {
		current := 3
		seqs := make([][]rune, 1, 10)
		for _, char := range code {
			move_possibilities := moveKeypad(current, char)
			curr_s := len(seqs)
			for s := 0; s < curr_s; s++ {
				for i, poss := range move_possibilities {
					if i == len(move_possibilities)-1 {
						seqs[s] = append(seqs[s], poss...)
						seqs[s] = append(seqs[s], 'A')
					} else {
						new_seq := make([]rune, len(seqs[s]))
						copy(new_seq, seqs[s])
						new_seq = append(new_seq, poss...)
						new_seq = append(new_seq, 'A')
						seqs = append(seqs, new_seq)
					}
				}
			}
			current = char
		}

		//for _, seq := range seqs {
		//	fmt.Println(string(seq))
		//}

		res := make([][]rune, 0, 100)

		for _, seq := range seqs {
			res = append(res, solveDirpad('A', seq, make([][]rune, 1, 100))...)
		}

		res2 := make([][]rune, 0, 100)
		for _, seq := range res {
			sols := solveDirpad('A', seq, make([][]rune, 1))
			res2 = append(res2, sols...)
		}

		min_steps := res2[0]
		for _, r := range res2 {
			if len(min_steps) > len(r) {
				min_steps = r
			}
		}
		acc += len(min_steps) * code_num[i]
	}
	fmt.Println(acc)
}

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

func solveDirpad(curr rune, seq []rune, res [][]rune) [][]rune {
	if len(seq) == 0 {
		return res
	}

	target := seq[0]
	next_moves := moveDirpad(curr, target)

	res_size := len(res)
	for s := 0; s < res_size; s++ {
		for i, poss := range next_moves {
			if i == len(next_moves)-1 {
				res[s] = append(res[s], poss...)
			} else {
				new_seq := make([]rune, len(res[s]))
				copy(new_seq, res[s])
				new_seq = append(new_seq, poss...)
				res = append(res, new_seq)
			}
		}
	}
	for s, _ := range res {
		res[s] = append(res[s], 'A')
	}
	return solveDirpad(target, seq[1:], res)
}

func moveKeypad(pos int, target int) [][]rune {
	col := roundUp(target, 3) - roundUp(pos, 3)
	row := target - pos - (col * 3)

	return rowColSeq(row, col, pos, target, 1, false)
}

func moveDirpad(pos rune, target rune) [][]rune {
	if pos == target {
		return make([][]rune, 0)
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
	return rowColSeq(row, col, dir_pad[pos], dir_pad[target], 4, true)
}

func rowColSeq(row, col, pos, target, blocked int, rev bool) [][]rune {
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
			return [][]rune{append(seq1, seq2...)}
		} else {
			return [][]rune{append(seq2, seq1...)}
		}
	}

	if len(seq1) == 0 || len(seq2) == 0 {
		return [][]rune{append(seq1, seq2...)}
	}

	return [][]rune{append(seq1, seq2...), append(seq2, seq1...)}
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
