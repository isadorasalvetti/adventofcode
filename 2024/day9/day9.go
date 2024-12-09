package main

import (
	"fmt"
	"os"
)

func main() {
	readDisk("../_sample/day9.txt")
}

func readDisk(file string) {
	contents_, _ := os.ReadFile(file)
	contents := string(contents_)
	lines, gaps, file_dists := split(contents)
	filed, unfiled := organizeFiles2(lines, gaps, file_dists, make([][]int, 0), make(map[int]int))
	fmt.Println(filed, unfiled)
	acc := 0
	places := 0
	for _, file := range filed {
		places += unfiled[places]
		for i := 0; i < file[1]; i++ {
			acc += file[0] * (places + i)
			fmt.Println(file[0], places+i)
		}
		places += file[1]
	}
	fmt.Println(acc)
}

func split(line string) ([]int, []int, []int) {
	files := make([]int, len(line)/2)
	file_dists := make([]int, len(line)/2)
	gaps := make([]int, len(line)/2)
	dist_acc := 0
	for i, val := range line {
		if i%2 == 0 {
			files[i/2] = int(val - '0')
			file_dists[i/2] = dist_acc
			dist_acc += files[i/2]
		} else {
			gaps[i/2] = int(val - '0')
			dist_acc += gaps[i/2]
		}
	}
	return files, gaps, file_dists
}

func organizeFiles(files, gaps, file_dists []int, filed_gaps [][]int, unfiled_gaps map[int]int) ([][]int, map[int]int) {
	for i := 0; i < len(files); i++ {
		if files[i] > 0 {
			filed_gaps = append(filed_gaps, []int{i, files[i]})
			files[i] = 0
		}
		if gaps[i] > 0 {
			for j := len(files) - 1; j >= 0; j-- {
				if j == 0 {
					filed_gaps = append(filed_gaps, []int{0, gaps[i]})
					gaps[i] = 0
				} else if files[j] > 0 && files[j] <= gaps[i] {
					filed_gaps = append(filed_gaps, []int{j, files[j]})
					unfiled_gaps[file_dists[j]] = files[j]
					gaps[i] = gaps[i] - files[j]
					files[j] = 0
					return organizeFiles(files, gaps, file_dists, filed_gaps, unfiled_gaps)
				}
			}
		}
	}
	return filed_gaps, unfiled_gaps
}

func organizeFiles2(files, gaps, file_dists []int, filed_gaps [][]int, unfiled_gaps map[int]int) ([][]int, map[int]int) {
	for f := 0; f < len(files)*2-1; f++ {
		if f%2 == 0 {
			if files[f/2] > 0 {
				fmt.Println(f/2, files[f/2])
				filed_gaps = append(filed_gaps, []int{f, files[f/2]})
				files[f/2] = 0
			}
		} else {
			for j := len(files) - 1; j >= f/2; j-- {
				for i := 0; i < len(gaps); i++ {
					if j == 0 {
						fmt.Println(j, files[j])
						filed_gaps = append(filed_gaps, []int{i, gaps[i]})
						files[j] = 0
					} else if files[j] > 0 && files[j] <= gaps[i] {
						fmt.Println(j, files[j])
						filed_gaps = append(filed_gaps, []int{j, files[j]})
						unfiled_gaps[file_dists[j]] = files[j]
						gaps[i] = gaps[i] - files[j]
						files[j] = 0
						return organizeFiles2(files, gaps, file_dists, filed_gaps, unfiled_gaps)
					}
				}
			}
		}
	}
	return filed_gaps, unfiled_gaps
}

func reReadDisk(toRead string, left_file_id, right_file_id, right_file_size, lenght_remaining, block_position int, is_file bool, acc int) int {
	if len(toRead) == 0 {
		for right_file_size > 0 {
			acc += right_file_id * block_position
			block_position += 1
			right_file_size -= 1
		}
		return acc
	}

	currLen := int(toRead[0] - '0')
	if is_file {
		for i := 0; i < currLen; i++ {
			acc += left_file_id * block_position
			block_position += 1
		}
		left_file_id += 1
		return reReadDisk(toRead[1:], left_file_id, right_file_id, right_file_size, lenght_remaining, block_position, false, acc)
	}

	lenght_remaining = currLen
	toRead = toRead[1:]
	for lenght_remaining > 0 {
		if right_file_size == 0 {
			if len(toRead) < 2 {
				return acc
			}
			right_file_size = int(toRead[len(toRead)-2] - '0')
			right_file_id -= 1
			toRead = toRead[:len(toRead)-2]
		}
		for right_file_size > 0 && lenght_remaining > 0 {
			acc += right_file_id * block_position
			block_position += 1
			lenght_remaining -= 1
			right_file_size -= 1
		}
	}
	return reReadDisk(toRead, left_file_id, right_file_id, right_file_size, lenght_remaining, block_position, true, acc)
}

func deFragment(toRead string, left_file_id, right_file_id, block_position_left, block_position_right int, is_file bool, acc int) int {
	if right_file_id <= left_file_id {
		return acc
	}

	fmt.Println(toRead)

	currLen := int(toRead[0] - '0')
	if is_file {
		for i := 0; i < currLen; i++ {
			fmt.Println(left_file_id, block_position_left, "left")
			acc += left_file_id * block_position_left
			block_position_left += 1
		}
		left_file_id += 1
		return deFragment(toRead[1:], left_file_id, right_file_id, block_position_left, block_position_right, false, acc)
	}

	right_pos := len(toRead) - 2
	to_move := int(toRead[right_pos] - '0')
	space_id := 0
	skipped_spaces := 0
	fmt.Println("Moving", to_move)
	for space_id <= len(toRead)-2 {
		space_to_fill := int(toRead[space_id] - '0')
		fmt.Println("Moving into ", space_id, space_to_fill)

		if space_to_fill >= to_move {
			for to_move > 0 {
				fmt.Println(right_file_id, (block_position_left + skipped_spaces), "right moves")
				acc += right_file_id * (block_position_left + skipped_spaces)
				block_position_right -= 1
				space_to_fill -= 1
				to_move -= 1
			}
			if space_to_fill > 0 {
				fmt.Println("Add new space")
				toRead = toRead[0:space_id] + string('0'+space_to_fill) + toRead[space_id+1:len(toRead)-2]
				is_file = false
			} else {
				toRead = toRead[0:space_id] + toRead[space_id+1:len(toRead)-2]
				is_file = true
			}
			right_file_id -= 1
			break
		}
		skipped_spaces += int(toRead[space_id] - '0')
		skipped_spaces += int(toRead[space_id+1] - '0')
		space_id += 2
	}
	if to_move > 0 {
		for to_move > 0 {
			fmt.Println(right_file_id, block_position_right, "right stays")
			acc += right_file_id * block_position_right
			block_position_right -= 1
			to_move -= 1
		}
		is_file = false
		block_position_right -= int(toRead[(len(toRead)-2)] - '0')
		toRead = toRead[:len(toRead)-2]
		right_file_id -= 1
	}
	return deFragment(toRead, left_file_id, right_file_id, block_position_left, block_position_right, is_file, acc)
}
