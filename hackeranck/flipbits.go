package main

import "fmt"

/*
	This one is funny
 	convert n in base2
	flip the bits and convert it back in base10

*/

func flippingBits(n int64) int64 {

	var answer int64
	res := n
	// base 2 integer stored with 32 bits
	baseTwo := make([]int, 32)

	// tant que le resultat de la division est > 0
	// on stock le reste en mode "little endian" ce qui donne le résultat en base2
	i := 31
	for res > 0 {
		if res%2 == 0 {
			baseTwo[i] = 0
		} else {
			baseTwo[i] = 1
		}
		res = res / 2
		i--

	}

	// flip it
	for i := range baseTwo {
		if baseTwo[i] == 0 {
			baseTwo[i] = 1
		} else {
			baseTwo[i] = 0
		}
	}

	// back to base10
	j := 0
	for i := len(baseTwo) - 1; i >= 0; i-- {
		if baseTwo[i] == 1 {
			answer += int64(pow(2, j))
		}
		j++
	}

	return answer

}

func pow(n, e int) int {
	if e == 0 {
		return 1
	}
	r := n
	for i := 1; i < e; i++ {
		r *= n
	}
	return r
}

func main() {
	var input int64 = 1
	// 00000000000000000000000000000001
	// expecting 4294967294
	fmt.Println(flippingBits(input))

}
