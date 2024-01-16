package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gnoale/adventofcode/puzzlein"
)

type adjacency map[string]string

type seatmap [][]string

func buildSeatmap(in []string) seatmap {
	s := seatmap{}
	for _, l := range in {
		s = append(s, strings.Split(l, ""))
	}
	return s
}

func buildAdjacency(s seatmap) adjacency {
	adj := adjacency{}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			x := strconv.Itoa(j)
			y := strconv.Itoa(i)
			adj[fmt.Sprintf("%v:%v", x, y)] = s[i][j]
		}
	}
	return adj
}

func (a adjacency) getPositions(x, y int) []string {

	x1 := strconv.Itoa(x - 1)
	x2 := strconv.Itoa(x + 1)
	y1 := strconv.Itoa(y - 1)
	y2 := strconv.Itoa(y + 1)

	pos := make([]string, 0)
	pos = append(pos, fmt.Sprintf("%v:%v", x1, y))
	pos = append(pos, fmt.Sprintf("%v:%v", x2, y))
	pos = append(pos, fmt.Sprintf("%v:%v", x1, y1))
	pos = append(pos, fmt.Sprintf("%v:%v", x1, y2))
	pos = append(pos, fmt.Sprintf("%v:%v", x2, y1))
	pos = append(pos, fmt.Sprintf("%v:%v", x2, y2))
	pos = append(pos, fmt.Sprintf("%v:%v", x, y1))
	pos = append(pos, fmt.Sprintf("%v:%v", x, y2))

	return pos
}

// if all adjacent are empty, return occupied, else stays empty
func (a adjacency) staysEmpty(x, y int) string {
	pos := a.getPositions(x, y)
	for _, p := range pos {
		if a[p] == "#" {
			return "L"
		}
	}
	return "#"
}

// if at least 4 adjacent are occupied, return empty
func (a adjacency) staysOccupied(x, y int) string {
	pos := a.getPositions(x, y)
	c := 0
	for _, p := range pos {
		if a[p] == "#" {
			c++
		}
	}
	if c >= 4 {
		return "L"
	}
	return "#"
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// if current == previous : count seats and return
func (s seatmap) run(a adjacency, news seatmap) int {
	equal := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		t := make([]string, len(s[i]))
		copy(t, s[i])
		news = append(news, t)
		//fmt.Println(news[i])
		//fmt.Println(s[i])
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == "L" {
				news[i][j] = a.staysEmpty(j, i)
			} else if s[i][j] == "#" {
				news[i][j] = a.staysOccupied(j, i)
			}
		}
		//fmt.Println(s[i], news[i])
		equal[i] = Equal(s[i], news[i])
	}
	// new round ?
	a = buildAdjacency(news)
	for _, e := range equal {
		if !e {
			n := seatmap{}
			return news.run(a, n)
		}
	}
	// end
	oc := 0
	for _, row := range news {
		for _, seat := range row {
			if seat == "#" {
				oc++
			}
		}
	}
	return oc
}

func main() {
	in, err := puzzlein.GetStr("./day11/input")
	if err != nil {
		panic(err)
	}

	s := buildSeatmap(in)
	adj := buildAdjacency(s)

	news := seatmap{}
	res := s.run(adj, news)
	fmt.Println(res)

}
