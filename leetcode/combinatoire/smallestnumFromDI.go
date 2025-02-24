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
