package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gnoale/adventofcode/puzzlein"
)

type Stack []vertex

func (s Stack) Push(v vertex) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, vertex) {
	l := len(s)
	return s[:l-1], s[l-1]
}

type graph map[string]vertex

type vertex struct {
	color    string
	count    int
	visiting bool
	visited  bool
	child    []vertex
}

func (g graph) populate(bags []string) {
	for _, line := range bags {
		v := vertex{}

		sl := strings.Split(line, ",") // sl = ["muted white bags contain 4 dark orange bags" "3 bright white bags."]
		s := strings.Split(sl[0], " ") // s = ["muted" "white" "bags" "contain" ...]

		v.color = strings.Join(s[:2], " ") // "muted white"
		v.count = 1

		color := strings.Join(s[5:7], " ")

		if color != "other bags." {
			v1 := vertex{}
			v1.color = color                 // "dark orange"
			v1.count, _ = strconv.Atoi(s[4]) // 4
			v.child = []vertex{v1}

			for i := 1; i < len(sl); i++ {
				vn := vertex{}
				s = strings.Split(sl[i], " ") // s = ["3", "bright", "white"...]
				vn.color = strings.Join(s[2:4], " ")
				vn.count, _ = strconv.Atoi(s[1])
				v.child = append(v.child, vn)
			}
			g[v.color] = v
		}
	}
}

func (g graph) topologicalSort() []string {
	stack := Stack{}
	for _, v := range g {
		if !v.visited {
			if !g.dfs(v, stack) {
				return []string{}
			}
		}
	}
	res := make([]string, 0)
	var v vertex
	for len(stack) > 0 {
		stack, v = stack.Pop()
		res = append(res, v.color)
	}
	return res

}

func (g graph) dfs(v vertex, s Stack) bool {
	v.visiting = true
	for _, cv := range v.child {
		cv = g[cv.color]
		if !cv.visited {
			if cv.visiting {
				return false
			}
			childres := g.dfs(cv, s)
			if !childres {
				return false
			}
		}
	}
	s.Push(v)
	v.visited = true
	v.visiting = false
	return true

}

func main() {

	b, err := puzzlein.GetStr("./day7/inputest")
	if err != nil {
		panic(err)
	}

	// part1

	// part2
	g := graph{}
	g.populate(b)
	//fmt.Println(g)

	s := g.topologicalSort()
	fmt.Println(s)

}
