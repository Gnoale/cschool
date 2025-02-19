package leetcode75

/*
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

n2 par défaut

start backward ?
	Non ça change rien

if i == 0
	output[i] = output[i+1] * output[i+2] * ouput[i+n]

if i > 0
	output[i] = output

Sans utiliser la division...
Avec : calculer le produit, puis pour chaque position on divise le total par la valeur de output[i]

Utiliser 2 pointeurs pour ne traverser que 2 fois l'array



*/

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	left := 1
	for i := 0; i < len(nums); i++ {
		output[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = right * output[i]
		right *= nums[i]
	}

	return output
}

/*
# Intuition

$$\prod_{i=0,i\neq i}^{n-1} p(i) = p(0) * ... *p(i-1) * p(i+1)* ... * p(n-1)$$

# Approach
<!-- Describe your approach to solving the problem. -->
Solving it first from left to right
$$ l(i) = \prod_{i=0}^{i-1} $$

And from right to left
$$p(i) = \prod_{i=l(n-1)}^{i+1}$$


# Complexity
- Time complexity:

O(n)
- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
O(n)
# Code
```
func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	left := 1
	for i := 0; i < len(nums); i++ {
		output[i] = left
		left *= nums[i]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = right * output[i]
		right *= nums[i]
	}
	return output
}
```


*/
