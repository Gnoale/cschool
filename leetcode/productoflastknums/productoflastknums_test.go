package productoflastknums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductOfNumbers(t *testing.T) {
	productOfNumbers := Constructor()
	productOfNumbers.Add(3)
	productOfNumbers.Add(0)
	productOfNumbers.Add(2)
	productOfNumbers.Add(5)
	productOfNumbers.Add(4)
	assert.Equal(t, 20, productOfNumbers.GetProduct(2))
	assert.Equal(t, 40, productOfNumbers.GetProduct(3))
	assert.Equal(t, 0, productOfNumbers.GetProduct(4))
	productOfNumbers.Add(8)
	assert.Equal(t, 32, productOfNumbers.GetProduct(2))
}
