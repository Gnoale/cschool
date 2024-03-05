package main

import (
	"fmt"
)

// génération de combinaisons pour le fun
// fonctionne bien avec une fenêtre glissante
// impossible à setup correctement avec des pointeurs pour gérer les possibles valeurs de k
func combine(s []int, k int) [][]int {
	combo := [][]int{}
	// s[i] is our head
	for i := 0; i <= len(s)-k; i++ {
		for j := i + 1; j < len(s); j++ {
			if k+j-1 > len(s) {
				break
			}
			// this is the tail s[j:j+k-1]
			w := []int{s[i]}
			w = append(w, s[j:j+k-1]...)
			combo = append(combo, w)
		}
	}
	return combo
}

func birthday(s []int32, d int32, m int32) int32 {
	/* il faut uniquement regarder les segments consecutifs
	   pas besoin de faire de combinaisons
	   chaque element est utilisable 2  fois.
	   ce n'est pas possible de le faire avec des pointeurs
	   il faut utiliser une fenetre glissante
	*/
	var result int32
	for i := 0; i <= len(s)-int(m); i++ {
		var sum int32
		if int(m)+i-1 > len(s) {
			break
		}
		for _, v := range s[i : i+int(m)] {
			sum += v
		}
		if sum == d {
			result++
		}
	}
	return int32(result)
}

func main() {

	fmt.Println(combine([]int{1, 2, 3, 4, 5}, 2))

	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))

}
