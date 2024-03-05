package main

import (
	"fmt"
	"sort"
)

func findMax(category []int, price []int) int {

	products := []product{}

	for i, cat := range category {
		products = append(products, product{
			category: cat,
			price:    price[i],
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].price < products[j].price
	})

	seen := map[int]bool{}
	i := 0
	total := 0

	// iterate over our products sorted by ascending order over the price
	for _, product := range products {
		// if the product category has already been seen, we do not increment the rate
		if ok, _ := seen[product.category]; !ok {
			i += 1
			seen[product.category] = true
		}
		total += i * product.price
	}

	return total

}

type product struct {
	category int
	price    int
}

func main() {
	fmt.Println(findMax([]int{3, 1, 4, 1}, []int{12, 4, 8, 4}))
	/*
		trouver la succession optimale de produit en fonction du nombre de catégories visitées
		et du prix de l'item

		1*4 + 1*4 + 2*8 + 3*12 = 60

		l'idée est de trier les catégorie par prix
		et de trier les 2 en fonction du prix

		il faut merge les 2 liste en 1 seul objet

	*/
}
