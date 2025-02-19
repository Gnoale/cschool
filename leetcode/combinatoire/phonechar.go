package combinatoire

import (
	"strings"
)

var keyboard = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {

	results := []string{}
	if len(digits) == 0 {
		return results
	}
	dfs(digits, 0, "", &results)
	return results
}

func dfs(digits string, index int, path string, results *[]string) {
	if index == len(digits) {
		*results = append(*results, path)
		return
	}

	letters := getLetters(string(digits[index]))
	for _, letter := range letters {
		path = path + letter
		dfs(digits, index+1, path, results)
		// prune the last letter to backtrack
		path = path[:len(path)-1]
	}
}

func getLetters(digit string) []string {
	return strings.Split(keyboard[digit], "")
}
