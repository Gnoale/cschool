package aoc

import (
	"2023/pkg/utils"
	"bufio"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func Parts() {

	grid := [][]int{}

	i, j := 0, 0

	scanner := bufio.NewScanner(utils.ReadInput("d3.input"))
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		grid = append(grid, row)
		j = 0

		for len(line) > 0 {
			//runtime.Breakpoint()
			r, size := utf8.DecodeRuneInString(line)

			// find numbers
			currentInt := []int{}
			// while we have a digit move forward and record it
			for len(line) > 0 {
				n, err := strconv.Atoi(string(r))
				if err != nil {
					// record gear
					if r == '*' {
						row[j] = -2
					} else if r != '.' {
						row[j] = -1
					}
					break
				}
				currentInt = append(currentInt, n)
				line = line[size:]
				j += 1
				r, size = utf8.DecodeRuneInString(line)
			}

			// build the final current number
			var number int
			if len(currentInt) > 0 {
				for i := len(currentInt) - 1; i >= 0; i-- {
					number += currentInt[i] * utils.Pow(10, len(currentInt)-1-i)
				}
			}
			if number > 0 {
				for k := j - len(currentInt); k < j; k++ {
					row[k] = number
				}
				//fmt.Println(number)
			}

			line = line[size:]
			j += 1
		}
		i += 1
	}

	enginePart2(grid)
}

/*
	func part1(grid [][]int) {
		total := 0
		p := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				n := grid[i][j]
				found := false
				// this mean we already found a match for the previous check
				// avoid eventually finding another match and add it up several times
				if n == p {
					found = true
				}
				//runtime.Breakpoint()
				// check in all direction if -1 appears
				// make sure to add it one time only

				// check previous row from j-1 to j+1
				// . . . ! . = [i-1][j]
				// 3 3 3 . . = [i][j]
				if n > 0 && !found {
					if i > 0 {
						k := 0
						if j > 1 {
							k = j - 1
						}
						for ; k <= j+1 && k < len(grid[i-1]); k++ {
							if grid[i-1][k] == -1 {
								total += n
								p = n
								found = true
								break
							}
						}
					}
					if !found {
						// check right
						if j+1 < len(grid[j]) {
							if grid[i][j+1] == -1 {
								total += n
								p = n
								found = true
							}
						}
					}
					if !found {
						// check left
						if j > 0 {
							if grid[i][j-1] == -1 {
								total += n
								p = n
								found = true
							}
						}
					}
					if !found {
						// check next row
						if i+1 < len(grid) {
							k := 0
							if j > 1 {
								k = j - 1
							}
							for ; k <= j+1 && k < len(grid[i+1]); k++ {
								if grid[i+1][k] == -1 {
									total += n
									p = n
									break
								}
							}
						}
					}
				}

				// reset previous
				if n == 0 {
					p = 0
				}
			}
			fmt.Println(grid[i], total)
		}
	}
*/
func enginePart2(grid [][]int) {
	total := 0
	//runtime.Breakpoint()
	factors := [][]int{}
	for i := 0; i < len(grid); i++ {
		// factor n numbers around the gear
		if len(factors) > 0 {
			for _, factor := range factors {
				if len(factor) >= 2 {
					for i := 0; i < len(factor); i++ {
						for j := i + 1; j < len(factor); j++ {
							total += factor[i] * factor[j]
						}
					}
				}
			}

		}
		factors = [][]int{}
		for j := 0; j < len(grid[i]); j++ {
			n := grid[i][j]
			// found a gear *
			if n == -2 {
				f := []int{}
				// check previous row from j-1 to j+1
				// . . . ! . = [i-1][j]
				// 3 3 3 . . = [i][j]
				if i > 0 {
					k := 0
					if j > 1 {
						k = j - 1
					}
					/*
						..7.3.....
						..6*2.....

						this should lead to
							factors = []int{7,3,6,2}
							7*3 + 7*6 + 7*2 + 3*6 + 3*2 + 6*2

						I need to record the previous numner to avoid adding it twice
					*/
					p := 0
					for ; k <= j+1 && k < len(grid[i-1]); k++ {
						num := grid[i-1][k]
						// initial position
						if num > 0 && (k == 0 || k == j-1) {
							f = append(f, num)
							p = num
						}
						if num > 0 && p != num {
							f = append(f, num)
							p = num
						}
					}
				}
				// check right
				if j+1 < len(grid[j]) && grid[i][j+1] > 0 {
					f = append(f, grid[i][j+1])
				}
				// check left
				if j > 0 && grid[i][j-1] > 0 {
					f = append(f, grid[i][j-1])
				}
				// check next row
				if i+1 < len(grid) {
					k := 0
					if j > 1 {
						k = j - 1
					}
					p := 0
					for ; k <= j+1 && k < len(grid[i+1]); k++ {
						num := grid[i+1][k]
						// initial position
						if num > 0 && (k == 0 || k == j-1) {
							f = append(f, num)
							p = num
						}
						if num > 0 && p != num {
							f = append(f, num)
							p = num
						}
					}
				}
				factors = append(factors, f)
			}
		}
	}
	fmt.Println(total)
}
