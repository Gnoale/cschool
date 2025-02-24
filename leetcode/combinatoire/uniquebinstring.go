package combinatoire

func findDifferentBinaryString(nums []string) string {

	mapNum := map[string]bool{}
	for _, num := range nums {
		mapNum[num] = true
	}

	//arrangements := []string{}

	// k arragement of n in numbers with repetitions
	var permute func(numbers string, permutation string, k int) string
	permute = func(numbers string, permutation string, k int) string {
		// base condition
		if len(permutation) == k {
			_, ok := mapNum[permutation]
			if !ok {
				return permutation
			}
			return ""
		}
		// find next node
		for _, number := range numbers {
			permutation += string(number)
			result := permute(numbers, permutation, k)
			if result != "" {
				return result
			}
			// backtrack
			permutation = permutation[:len(permutation)-1]
		}
		return ""
	}

	return permute("01", "", len(nums))

}
