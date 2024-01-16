package main

import (
	"fmt"
	"math/rand"
)

func merge_sort(A []int) []int {

	N := len(A)

	if N == 1 {
		fmt.Printf("Leaf %v\n", A)
		return A
	}

	L := A[:N/2]
	R := A[N/2:]
	fmt.Printf("Left tree %v", L)
	fmt.Printf(" Right tree %v\n", R)
	L = merge_sort(L)
	R = merge_sort(R)
	fmt.Printf("Left tree %v", L)
	fmt.Printf(" Right tree %v\n", R)
	fmt.Printf("Merge %v and %v\n", L, R)

	return merge(L, R)

}

func merge(L []int, R []int) []int {

	m := make([]int, 0)

	i, j := 0, 0
	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			m = append(m, L[i])
			i++
		} else {
			m = append(m, R[j])
			j++
		}
	}
	for ; i < len(L); i++ {
		m = append(m, L[i])
	}
	for ; j < len(R); j++ {
		m = append(m, R[j])
	}

	fmt.Printf("Merged array %v\n\n", m)
	return m
}

func main() {

	A := make([]int, 20)
	for i, _ := range A {
		A[i] = rand.Intn(20)
	}
	fmt.Println(A)
	B := merge_sort(A)
	fmt.Println(B)

}
