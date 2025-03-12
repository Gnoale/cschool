package combinatoire

import "sort"

func smallestNumber(pattern string) string {

	permutations := []string{}

	numbers := "123456789"
	visited := make(map[rune]bool)

	uniquePermute(numbers, "", pattern, &permutations, visited, len(pattern)+1)

	sort.Strings(permutations)
	return permutations[0]

}

func uniquePermute(elements string, permutation string, pattern string, results *[]string, visited map[rune]bool, k int) {
	// base condition
	if len(permutation) == k {
		*results = append(*results, permutation)
		return
	}
	// find next node
	for _, element := range elements {
		if visited[element] {
			continue
		}
		if len(permutation) > 0 {
			if pattern[len(permutation)-1] == 'I' {
				if rune(permutation[len(permutation)-1]) > element {
					continue
				}
			} else {
				if rune(permutation[len(permutation)-1]) < element {
					continue
				}
			}
		}
		visited[element] = true
		permutation += string(element)
		uniquePermute(elements, permutation, pattern, results, visited, k)
		// backtrack
		permutation = permutation[:len(permutation)-1]
		visited[element] = false
	}
}
