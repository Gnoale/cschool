/*

Implementation of a heap

A heap is a kind of complete binary tree

The tree is either a max or min heap

In a max heap, the parent is greater than its children
In a min heap, the parent is smaller than its children

Super usefull to keep a datastructure sorted with O(log n) complexity
*/

package heap

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Heap struct {
	Nodes     []*Node
	HeapSize  int
	LastIndex int
}

func NewHeap(size int) *Heap {
	return &Heap{
		Nodes:     make([]*Node, size),
		HeapSize:  size,
		LastIndex: -1,
	}
}

func (h *Heap) Min() int {
	return h.Nodes[0].Value
}

func (h *Heap) Add(item int) {
	if h.LastIndex > h.HeapSize {
		panic("Heap is full")
	}
	h.LastIndex++
	h.Nodes[h.LastIndex] = &Node{Value: item}
	h.bubbleUp(h.LastIndex)
}

func (h *Heap) bubbleUp(index int) {
	parentIndex := (index - 1) / 2
	for index > 0 {
		if h.Nodes[parentIndex].Value > h.Nodes[index].Value {
			h.Nodes[parentIndex], h.Nodes[index] = h.Nodes[index], h.Nodes[parentIndex]
		}
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

// Pop the min value from the heap, arbitrary replace the root with the last element in the heap
// and then bubble down the new root to its correct position
func (h *Heap) Pop() int {
	min := h.Nodes[0]

	h.Nodes[0] = h.Nodes[h.LastIndex]

	// pop the last element
	//h.Nodes = h.Nodes[:h.LastIndex]
	h.LastIndex--
	h.bubbleDown(0)

	return min.Value
}

func (h *Heap) bubbleDown(itemPosition int) {
	/*
		Cette étape est plus délicate
		on loop tant que l'un des enfants est plus petit que le parent
		s'il y a 2 enfants, on swap avec le plus petit
		ensuite on continue avec le sous arbre dans lequel se trouve notre valeur swappée

	*/
	index := 0
	for itemPosition < h.LastIndex/2+1 {
		item := h.Nodes[itemPosition]
		left := 2*itemPosition + 1
		right := 2*itemPosition + 2
		if right <= h.LastIndex && h.Nodes[right].Value < h.Nodes[left].Value {
			index = right
		} else {
			index = left
		}
		if item.Value > h.Nodes[index].Value {
			h.Nodes[itemPosition], h.Nodes[index] = h.Nodes[index], h.Nodes[itemPosition]
			itemPosition = index
		} else {
			break
		}
	}
}

func (h *Heap) Peek() int {
	return h.Nodes[0].Value
}
