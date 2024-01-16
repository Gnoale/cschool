package aoc

import (
	"2023/pkg/utils"
	"bufio"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/zyedidia/generic/stack"
)

type Node struct {
	V string
	L string
	R string
}

type Network struct {
	direction []int // LR[0,1] RL[1,0]
	i         int   // track current position in direction
	nodes     []Node
	graph     map[string]int // Node position in nodes
	mu        *sync.Mutex
}

func parse(scanner *bufio.Scanner) Network {
	nodes := []Node{}
	// parse direction
	dir := []int{}
	scanner.Scan()
	line := scanner.Text()
	for _, char := range line {
		if string(char) == "L" {
			dir = append(dir, 0)
		}
		if string(char) == "R" {
			dir = append(dir, 1)
		}
	}
	i := 0
	graph := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		node := Node{}
		left, right, _ := strings.Cut(line, " = ")
		node.V = left
		left, right, _ = strings.Cut(right, ", ")
		node.L = strings.Trim(left, "(")
		node.R = strings.Trim(right, ")")
		nodes = append(nodes, node)
		graph[node.V] = i
		i++
	}

	return Network{
		direction: dir,
		i:         0,
		nodes:     nodes,
		graph:     graph,
		mu:        &sync.Mutex{},
	}
}

func dfs(net Network, root Node) int {

	stack := stack.New[Node]()

	stack.Push(root)

	t := 0
	for stack.Size() > 0 {
		node := stack.Pop()
		if strings.HasSuffix(node.V, "Z") {
			break
		}
		t++

		// determine next node
		if net.i == len(net.direction) {
			net.i = 0
		}
		nextV := ""
		switch net.direction[net.i] {
		case 0:
			nextV = node.L
		case 1:
			nextV = node.R
		}
		//fmt.Println(node, nextV)
		next := net.nodes[net.graph[nextV]]
		stack.Push(next)
		net.i++
	}
	fmt.Println("steps", t)
	return t
}

func NavigateDesert(scanner *bufio.Scanner) {

	net := parse(scanner)

	root := net.nodes[net.graph["AAA"]]

	dfs(net, root)

	// part2
	roots := []Node{}
	for _, node := range net.nodes {
		if strings.HasSuffix(node.V, "A") {
			roots = append(roots, node)
		}
	}

	steps := []int{}
	for _, root := range roots {
		steps = append(steps, dfs(net, root))
	}
	fmt.Println(steps)

	// find ppcm (plus petit commun multiple
	// ppcm(x,y) = (x*y) / gcd(x,y)

	/*
		gcd Euclid algo

		gcd(123, 23)
		123 = 23 * 5 + 8
		23 = 8 * 2 + 7
		8 = 7 * 1 + 1 => gcd = 1

	*/
	sort.IntSlice.Sort(steps)

	n, m := steps[len(steps)-2], steps[len(steps)-1]
	gcd := utils.GcdRec(n, m)
	ppcm := n * m / gcd

	for i := len(steps) - 3; i >= 0; i-- {
		ppcm = (ppcm * steps[i]) / gcd
	}
	fmt.Println(ppcm)
}
