package main

import (
	"fmt"
	"sort"

	"github.com/Gnoale/adventofcode/graph"
	"github.com/Gnoale/adventofcode/puzzlein"
)

func countDiff(adapter []int) map[int]int {
	count := make(map[int]int)
	for i := 0; i < len(adapter); i++ {
		var diff int
		if i == 0 {
			diff = adapter[i] - 0
		} else {
			diff = adapter[i] - adapter[i-1]
		}
		if diff <= 3 {
			count[diff]++
		}
	}
	count[3]++
	return count
}

func isin(n int, child []int) bool {
	for _, v := range child {
		if v == n {
			return true
		}
	}
	return false
}

func buildGraph(adapter []int, graph map[int][]int, i, j int) {
	if i < len(adapter) {
		diff := adapter[i] - adapter[j]
		if diff < 3 {
			// add value and check other branch
			k := i + 1
			l := j
			// if the child isn't there, add it
			if !isin(adapter[i], graph[adapter[j]]) {
				graph[adapter[j]] = append(graph[adapter[j]], adapter[i])
			}
			buildGraph(adapter, graph, k, l)
		}
		if diff == 3 {
			// if the child isn't there, add it
			if !isin(adapter[i], graph[adapter[j]]) {
				graph[adapter[j]] = append(graph[adapter[j]], adapter[i])
			}
		}
		i++
		j++
		buildGraph(adapter, graph, i, j)
	}
}

func traverse(root int, graph map[int][]int, path *[]int, paths *int, level map[int]int) {
	// start
	if level[root] == 0 {
		*path = append(*path, root)
	}
	child := graph[root]
	// finish
	if len(child) == 0 {
		*paths++
		//fmt.Println(*path)
	}
	for _, node := range child {
		// for each branch copy value of current branch "path"
		// inside a new path "p"
		// then, fork with the new path
		p := make([]int, len(*path))
		p = *path
		p = append(p, node)
		//fmt.Println(node, p)
		traverse(node, graph, &p, paths, level)
	}
}

func main() {
	in, err := puzzlein.GetInt("input")
	if err != nil {
		panic(err)
	}
	sort.Ints(in)
	// p1
	//m := countDiff(in)
	//fmt.Println(m, m[3]*m[1])

	// p2
	// add the 0 at begining
	input := make([]int, len(in)+1)
	input[0] = 0
	for i := 0; i < len(in); i++ {
		input[i+1] = in[i]
	}

	// build a graph
	g := make(map[int][]int)
	buildGraph(input, g, 1, 0)
	// sort all childs (they are in reverse order) - useless ?
	//	for _, v := range g {
	//		sort.Ints(v)
	//	}

	fmt.Println(g)

	level := graph.Bfs(input[0], g)
	fmt.Println(level)

	path := make([]int, 0)
	var paths int
	traverse(input[0], g, &path, &paths, level)
	fmt.Println(paths)

}
