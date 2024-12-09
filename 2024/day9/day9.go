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
	fmt.Println(deFragment(contents, 0, len(contents)/2-1, 0, true, 0))
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

func deFragment(toRead string, left_file_id, righ_max_id, block_position int, is_file bool, acc int) int {
	if len(toRead) == 0 {
		return acc
	}

	fmt.Println(toRead)

	currLen := int(toRead[0] - '0')
	if is_file {
		for i := 0; i < currLen; i++ {
			fmt.Println(left_file_id, block_position, "left")
			acc += left_file_id * block_position
			block_position += 1
		}
		left_file_id += 1
		return deFragment(toRead[1:], left_file_id, righ_max_id, block_position, false, acc)
	}

	right_pos := len(toRead) - 2
	ri := righ_max_id
	fmt.Println("Looking for fit for ", currLen)
	for right_pos > 0 {
		file_to_place := int(toRead[right_pos] - '0')
		if file_to_place > 0 && file_to_place <= currLen {
			fp := int(toRead[right_pos] - '0')
			fmt.Println("Found", ri, file_to_place, currLen, "right")
			for fp > 0 {
				fmt.Println(ri, block_position, "right")
				acc += ri * block_position
				block_position += 1
				fp -= 1
			}
			diff := currLen - file_to_place
			toRead = toRead[:right_pos] + "0" + toRead[right_pos+1:]
			if diff > 0 && file_to_place > 0 {
				toRead = toRead[:1] + string('0'+diff) + toRead[1:]
				return deFragment(toRead[1:], left_file_id, righ_max_id, block_position, false, acc)
			}
			break
		}
		right_pos -= 2
		ri -= 1
	}
	return deFragment(toRead[1:], left_file_id, righ_max_id, block_position, true, acc)
}
