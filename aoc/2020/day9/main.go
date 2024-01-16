package main

import (
	"fmt"

	"github.com/Gnoale/adventofcode/puzzlein"
)

func getCombSum(num []int) map[int]bool {
	m := make(map[int]bool)
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {
			if i != j {
				m[num[j]+num[i]] = true
			}
		}
	}
	return m

}

func preCheck(xmas []int, p int) {
	j := 0
	var weak int
	serie := make([][]int, 0)
	for i := p; i < len(xmas); i++ {
		serie = append(serie, xmas[j:i])
		comb := getCombSum(serie[j])
		if _, in := comb[xmas[i]]; !in {
			weak = xmas[i]
			fmt.Println("weakness number", weak)
			findSumWeak(serie, weak)
		}
		j++
	}
}

func findSumWeak(serie [][]int, weak int) {
	for _, seq := range serie {
		sum := 0
		b := 0
		for i := b; i < len(seq); i++ {
			if sum > weak {
				sum = 0
				b = i
			} else if sum == weak {
				computeKey(seq[b:i])
				break
			} else {
				sum += seq[i]
			}
		}
	}
}

func computeKey(wserie []int) {
	min := 0
	max := 0
	for _, num := range wserie {
		if min == 0 {
			min = num
		} else if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Println(min, max, min+max)

}

func main() {
	in, err := puzzlein.GetInt("./day9/input")
	if err != nil {
		panic(err)
	}
	preCheck(in, 25)

}
