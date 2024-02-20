package main

import "fmt"

/*
	trouver les entiers présents dans une seule liste
	2 maps ?
*/

func findDifference(nums1 []int, nums2 []int) [][]int {

	diff := make([][]int, 2)
	m1 := map[int]any{}
	m2 := map[int]any{}

	for _, num := range nums1 {
		m1[num] = true
	}

	for _, num := range nums2 {
		m2[num] = true
	}

	s := []int{}
	for num, _ := range m1 {
		_, ok := m2[num]
		if !ok {
			s = append(s, num)
		}
	}
	diff[0] = s

	s = []int{}
	for num, _ := range m2 {
		_, ok := m1[num]
		if !ok {
			s = append(s, num)
		}
	}
	diff[1] = s

	return diff

}

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}
