package graph

type Stack []int

func (s Stack) Push(v int) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func Bfs(root int, graph map[int][]int) map[int]int {
	s := make(Stack, 0)
	s = s.Push(root)
	visited := make(map[int]int)
	visited[root] = 0
	for len(s) > 0 {
		var toVisit int
		s, toVisit = s.Pop()
		next := graph[toVisit]
		for _, n := range next {
			if _, in := visited[n]; !in {
				s = s.Push(n)
				visited[n] = visited[toVisit] + 1
			}
		}
	}

	return visited
}
