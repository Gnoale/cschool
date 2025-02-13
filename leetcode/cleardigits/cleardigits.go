package cleardigits

import "fmt"

func findDigit(s string) int {
	for i, c := range s {
		if c >= '0' && c <= '9' {
			return i
		}
	}
	return -1
}

func findChar(index int, s string) int {
	for i := index; i >= 0; i-- {
		if s[i] >= 'a' && s[i] <= 'z' {
			return i
		}
	}
	return -1
}

func clearDigits(s string) string {

	for s != "" {

		firstDigitPosition := findDigit(s)
		fmt.Println(firstDigitPosition)
		if firstDigitPosition == -1 {
			return s
		}

		charPosition := findChar(firstDigitPosition, s)
		if charPosition == -1 {
			return s
		}

		s = s[:charPosition] + s[firstDigitPosition+1:]
	}
	return s
}
