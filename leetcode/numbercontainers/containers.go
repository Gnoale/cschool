package numbercontainers

import "fmt"

type NumberContainers struct {
	// number -> indexes
	numbersIndexes map[int][]int
	indexesNumbers map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		numbersIndexes: make(map[int][]int),
		indexesNumbers: make(map[int]int),
	}
}

func binarySearch(indexes []int, index int) int {
	left, right := 0, len(indexes)-1
	for left <= right {
		mid := (left + right) / 2
		if indexes[mid] == index {
			return mid
		}
		if indexes[mid] < index {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func (this *NumberContainers) removeIndex(number int, index int) {
	indexes := this.numbersIndexes[number]
	pos := binarySearch(indexes, index)
	this.numbersIndexes[number] = append(indexes[:pos], indexes[pos+1:]...)
	if len(this.numbersIndexes[number]) == 0 {
		delete(this.numbersIndexes, number)
	}
}

func (this *NumberContainers) addIndex(number int, index int) {
	indexes, ok := this.numbersIndexes[number]
	if !ok {
		this.numbersIndexes[number] = []int{index}
		return
	}

	if index > indexes[len(indexes)-1] {
		this.numbersIndexes[number] = append(indexes, index)
		return
	}

	if index < indexes[0] {
		this.numbersIndexes[number] = append([]int{index}, indexes...)
		return
	}

	pos := binarySearch(indexes, index)
	fmt.Println(indexes)
	fmt.Println(pos, indexes[pos], index)
	if indexes[pos] == index {
		return
	}
	if indexes[pos] > index {
		this.numbersIndexes[number] = append(indexes[:pos], append([]int{index}, indexes[pos:]...)...)
		fmt.Println(this.numbersIndexes[number])
		return
	}

	panic("cannot add index")
}

func (this *NumberContainers) Change(index int, number int) {
	// check if index already has a number
	num, ok := this.indexesNumbers[index]
	if num == number {
		return
	}
	if ok {
		this.removeIndex(num, index)
	}
	this.addIndex(number, index)
	this.indexesNumbers[index] = number
}

func (this *NumberContainers) Find(number int) int {
	if indexes, ok := this.numbersIndexes[number]; ok {
		return indexes[0]
	}
	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
