package productoflastknums

type ProductOfNumbers struct {
	queue    []int
	products []int
	size     int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{queue: []int{}, products: []int{1}, size: 0}

}

/*
C'est une problème similaire à la "prefix sum"
Pour le résoudre de manière constante il faut stocker la valeur cumullée des produits
lorsque l'on ajoute un 0, on reset la liste des produits car le 0 annule tous les produits précédents
34ms vs 800ms avec la version naïve qui consiste à faire le produit des k derniers nums à chaque requête
O(1) vs O(k)
*/

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.size = 0
		this.products = []int{1}
	} else {
		this.size++
		this.products = append(this.products, this.products[this.size-1]*num)
	}
	this.queue = append(this.queue, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if this.size < k {
		return 0
	}
	return this.products[len(this.products)-1] / this.products[len(this.products)-1-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
