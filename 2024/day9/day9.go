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
	fmt.Println(reReadDisk(contents, 0, len(contents)/2, 0, 0, 0, true, 0))

	sum_pos := 0
	for _, i := range contents {
		sum_pos += int(i - '0')
	}
	fmt.Println(deFragment(contents, 0, len(contents)/2-1, 0, sum_pos, true, 0))
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
