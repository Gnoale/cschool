package combinatoire

func getHappyString(n int, k int) string {
	results := []string{}
	choices := []string{"a", "b", "c"}
	generateHappyString(n, "", &results, choices)
	if k > len(results) {
		return ""
	}
	//sort.Strings(results)
	return results[k-1]
}

func generateHappyString(n int, current string, results *[]string, choices []string) {
	if len(current) == n {
		*results = append(*results, current)
		return
	}
	for _, choice := range choices {
		if len(current) > 0 && string(current[len(current)-1]) == choice {
			continue
		}
		current += choice
		generateHappyString(n, current, results, choices)
		current = current[:len(current)-1]
	}
}
