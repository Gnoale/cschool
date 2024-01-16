package aoc

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/zyedidia/generic/stack"
)

func parsePipe(scanner *bufio.Scanner) [][]*tile {
	m := [][]*tile{}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []*tile{}
		for j, m := range strings.Split(line, "") {
			var t *tile
			switch m {
			case "|":
				t = northSouth()
			case "-":
				t = eastWest()
			case "L":
				t = northEast()
			case "J":
				t = northWest()
			case "7":
				t = southWest()
			case "F":
				t = southEast()
			case ".":
				t = groundTile()
			case "S":
				t = rootTile()
			}
			t.i = i
			t.j = j
			row = append(row, t)

		}
		m = append(m, row)
		i += 1
	}
	return m
}

// Think BEHAVIOR
type tile struct {
	src, dst string
	i, j     int
	depth    int
	sym      string
}

func (t *tile) connect(direction string) bool {
	return t.src == direction || t.dst == direction
}

func (t *tile) String() string {
	//return fmt.Sprintf("%s (%d,%d)", t.sym, t.i, t.j)
	return t.sym
}

// enum tiles
// use a constructor to get a distinct object (with its own reference)
// for tracking them across the dfs

func northSouth() *tile {
	return &tile{src: "north", dst: "south", sym: "|"} // |
}

func eastWest() *tile {
	return &tile{src: "east", dst: "west", sym: "-"} // -
}

func northEast() *tile {
	return &tile{src: "north", dst: "east", sym: "L"} // L
}
func northWest() *tile {
	return &tile{src: "north", dst: "west", sym: "J"} // J
}

func southWest() *tile {
	return &tile{src: "south", dst: "west", sym: "7"} // 7
}

func southEast() *tile {
	return &tile{src: "south", dst: "east", sym: "F"} // F
}

func rootTile() *tile {
	return &tile{src: "any", dst: "any", sym: "S"} // S
}

func groundTile() *tile {
	return &tile{src: "none", dst: "none", sym: "."} // .
}

func FindPaths(scanner *bufio.Scanner) {

	grid := parsePipe(scanner)

	root := &tile{}

	for _, row := range grid {
		for _, tile := range row {
			if tile.sym == "S" {
				root = tile
			}
		}
	}
	visited := visitTiles(grid, root)

	// part 2: compter les points in /  out
	// ray casting algo  cannot implement

	inside := false
	total := 0
	for _, row := range grid {
		previous := &tile{sym: "."}
		for _, tile := range row {
			if visited[tile] {
				switch tile.sym {
				case "|":
					inside = !inside
				case "-":
				case "L":
					previous = tile
				case "F":
					previous = tile
				case "7":
					if previous.sym == "L" {
						inside = !inside
					}
				case "J":
					if previous.sym == "F" {
						inside = !inside
					}
				case ".":
					panic(tile)
				}
			} else if inside {
				total += 1
				tile.sym = "%"
			}
		}
	}

	for _, row := range grid {
		for _, tile := range row {
			fmt.Print(tile.sym)
		}
		fmt.Println()
	}
	fmt.Printf("there %d tiles inside\n", total)

}

// to find adjacencies you need to check if the current pipe plugs the next one in any direction
func visitTiles(tiles [][]*tile, root *tile) map[*tile]bool {

	s := stack.New[*tile]()

	// track visited grid paths
	visited := make(map[*tile]bool)
	s.Push(root)
	max := 0

	for s.Size() > 0 {
		toVisit := s.Pop()
		visited[toVisit] = true
		if toVisit.depth > max {
			max = toVisit.depth
		}

		// now look in every directions if another tile connect
		if toVisit.i > 0 {
			next := tiles[toVisit.i-1][toVisit.j]
			if next.connect("south") && !visited[next] {
				next.depth = toVisit.depth + 1
				s.Push(next)
			}
		}
		if toVisit.i < len(tiles)-1 {
			next := tiles[toVisit.i+1][toVisit.j]
			if next.connect("north") && !visited[next] {
				next.depth = toVisit.depth + 1
				s.Push(next)
			}
		}
		if toVisit.j > 0 {
			//runtime.Breakpoint()
			next := tiles[toVisit.i][toVisit.j-1]
			if next.connect("east") && !visited[next] {
				next.depth = toVisit.depth + 1
				s.Push(next)
			}
		}
		if toVisit.j < len(tiles[toVisit.i])-1 {
			next := tiles[toVisit.i][toVisit.j+1]
			if next.connect("west") && !visited[next] {
				next.depth = toVisit.depth + 1
				s.Push(next)
			}
		}
	}

	fmt.Println(max)
	return visited
}
