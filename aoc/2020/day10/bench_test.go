package main

import (
	"sort"
	"testing"

	"github.com/Gnoale/adventofcode/graph"
	"github.com/Gnoale/adventofcode/puzzlein"
)

var input []int
var g map[int][]int
var level map[int]int

func init() {
	in, err := puzzlein.GetInt("inputest")
	if err != nil {
		panic(err)
	}
	sort.Ints(in)
	input = make([]int, len(in)+1)
	input[0] = 0
	for i := 0; i < len(in); i++ {
		input[i+1] = in[i]
	}

	// build a graph
	g = make(map[int][]int)
	buildGraph(input, g, 1, 0)
	// sort all childs (they are in reverse order) - useless ?
	//	for _, v := range g {
	//		sort.Ints(v)
	//	}

	level = graph.Bfs(input[0], g)
}

func BenchmarkTraverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		path := make([]int, 0)
		//paths := make([][]int, 0)
		var paths int
		traverse(input[0], g, &path, &paths, level)

	}
}
