package leetcode75

func dailyTemperatures(temperatures []int) []int {
	response := make([]int, len(temperatures))
	stack := []int{}

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && (temperatures[stack[len(stack)-1]] < temperatures[i]) {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			response[j] = i - j
		}
		stack = append(stack, i)
	}

	return response
}
