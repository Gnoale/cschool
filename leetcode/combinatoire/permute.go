package combinatoire

// k arragement of n in numbers without repetitions
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

// k arragement of n in numbers with repetitions
func Permute(numbers string, permutation string, results *[]string, k int) {
	// base condition
	if len(permutation) == k {
		*results = append(*results, permutation)
		return
	}
	// find next node
	for _, number := range numbers {
		permutation += string(number)
		Permute(numbers, permutation, results, k)
		// backtrack
		permutation = permutation[:len(permutation)-1]
	}
}
