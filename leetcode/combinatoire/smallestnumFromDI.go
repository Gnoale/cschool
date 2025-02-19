package combinatoire

import "sort"

func smallestNumber(pattern string) string {

	permutations := []string{}

	numbers := "123456789"
	visited := make(map[rune]bool)
	UniquePermute(numbers, "", &permutations, visited, len(pattern)+1)

	candidates := []string{}
	for _, permutation := range permutations {
		ok := true
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'I' {
				if permutation[i] > permutation[i+1] {
					ok = false
					break
				}
			} else {
				if permutation[i] < permutation[i+1] {
					ok = false
					break
				}
			}
		}
		if ok {
			candidates = append(candidates, permutation)
		}
	}
	sort.Strings(candidates)
	return candidates[0]

}

func UniquePermute(numbers string, permutation string, results *[]string, visited map[rune]bool, k int) {
	// base condition
	if len(permutation) == k {
		*results = append(*results, permutation)
		return
	}
	// find next node
	for _, number := range numbers {
		if visited[number] {
			continue
		}
		visited[number] = true
		permutation += string(number)
		UniquePermute(numbers, permutation, results, visited, k)
		// backtrack
		permutation = permutation[:len(permutation)-1]
		visited[number] = false
	}

}
