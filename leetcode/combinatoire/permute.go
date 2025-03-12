package combinatoire

// k arragement of n elements without repetitions
func UniquePermute(elements string, permutation string, results *[]string, visited map[rune]bool, k int) {
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
		visited[element] = true
		permutation += string(element)
		UniquePermute(elements, permutation, results, visited, k)
		// backtrack
		permutation = permutation[:len(permutation)-1]
		visited[element] = false
	}
}

// k arragement of n elements with repetitions
func Permute(elements string, permutation string, results *[]string, k int) {
	// base condition
	if len(permutation) == k {
		*results = append(*results, permutation)
		return
	}
	// find next node
	for _, element := range elements {
		permutation += string(element)
		Permute(elements, permutation, results, k)
		// backtrack
		permutation = permutation[:len(permutation)-1]
	}
}
