package main

import "fmt"

/*
Avec 2 pointeurs, calculer la valeur de l'aire du rectangle
garder un maximum, déplacer le pointeur du segment le plus bas

*/

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		hl := height[left]
		hr := height[right]

		area := 0
		H := 0
		L := right - left

		if hl == hr {
			right -= 1
			left += 1
			H = hl
		} else if hl > hr {
			H = hr
			right -= 1
		} else {
			H = hl
			left += 1
		}
		area = H * L
		if area > max {
			max = area
		}

	}
	return max

}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}
