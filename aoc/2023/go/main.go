package main

import (
	"2023/pkg/aoc"
	"2023/pkg/utils"
	"bufio"
)

func main() {

	//aoc.Seed(bufio.NewScanner(utils.ReadInput("./inputs/d4.input")))
	//aoc.Race(bufio.NewScanner(utils.ReadInput("./inputs/d6.input")))
	//aoc.CamelPoker(bufio.NewScanner(utils.ReadInput("./inputs/d7.input")))
	//aoc.NavigateDesert(bufio.NewScanner(utils.ReadInput("./inputs/d8.input")))
	//aoc.FindSequences(bufio.NewScanner(utils.ReadInput("./inputs/d9.input")))
	aoc.FindPaths(bufio.NewScanner(utils.ReadInput("./inputs/d10.sample")))

}
