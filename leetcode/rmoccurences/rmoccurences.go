package rmoccurences

import "fmt"

func RemoveOccurrences(s string, part string) string {
	left := 0
	right := len(part)
	window := ""

	for right < len(s) {
		window = s[left:right]
		fmt.Println("window", window)

		if window == part {
			s = s[:left] + s[right:]
			if left > 0 {
				left = 0
				right = len(part)
			}
			fmt.Println("s", s)

		} else {
			left += 1
			right += 1
		}
	}
	if len(s) >= len(part) {
		window = s[len(s)-len(part):]

		if window == part {
			s = s[:len(s)-len(part)]
		}
	}
	return s

}
