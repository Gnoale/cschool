package leetcode75

import (
	"fmt"
	"strconv"
)

// we need to edit the array "in place" so we cannot afford to use the stack approach here
// also, the number of repetition must be encoded

/*

Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
func compress(chars []byte) int {
*/

func compress(chars []byte) int {
	// the approach consist in keeping 2 pointer
	// read and write
	// read will stop once the current value change
	// write will write the current char value into write
	// and the number of occurence (string encoded) at write + 1
	read, write := 0, 0
	for read < len(chars) {
		current := chars[read]
		count := 0
		// advance the read pointer until the char change
		for read < len(chars) && chars[read] == current {
			read++
			count++
		}
		fmt.Println(read, write)
		// write the number
		chars[write] = current
		write += 1
		if count > 1 {
			// store N as bytes in the array
			for _, char := range strconv.Itoa(count) {
				chars[write] = byte(char)
				write++
			}
		}
		// now the write pointer position must be <= the read one
	}
	fmt.Println(string(chars))
	return write

}
