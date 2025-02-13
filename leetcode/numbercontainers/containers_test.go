package numbercontainers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberContainers(t *testing.T) {

	operations := []string{"NumberContainers", "change", "find", "change", "find", "find", "find"}
	parameters := [][]int{{}, {1, 10}, {10}, {1, 20}, {10}, {20}, {30}}
	numberContainers := Constructor()
	result := []int{}
	for i, operation := range operations {
		if operation == "change" {
			numberContainers.Change(parameters[i][0], parameters[i][1])
		} else if operation == "find" {
			result = append(result, numberContainers.Find(parameters[i][0]))
		}
	}
	assert.Equal(t, []int{1, -1, 1, -1}, result)
}

func TestNumberContainers2(t *testing.T) {

	operations := []string{"NumberContainers", "change", "find", "change", "change", "change", "change", "change", "find", "change"}
	parameters := [][]int{{}, {1, 10}, {10}, {2, 10}, {3, 10}, {2, 20}, {7, 10}, {5, 10}, {10}, {5, 20}}
	numberContainers := Constructor()
	result := []int{}
	for i, operation := range operations {
		if operation == "change" {
			numberContainers.Change(parameters[i][0], parameters[i][1])
		} else if operation == "find" {
			result = append(result, numberContainers.Find(parameters[i][0]))
		}
	}
	assert.Equal(t, []int{1, 1}, result)
	fmt.Println(numberContainers.numbersIndexes)
}
