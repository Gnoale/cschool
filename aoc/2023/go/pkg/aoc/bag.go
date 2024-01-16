package aoc

import (
	"2023/pkg/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type pick struct {
	blue  int
	red   int
	green int
}

func Bag() {

	scanner := bufio.NewScanner(utils.ReadInput("d2.input"))

	// p1 count max
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	total := 0
	picks := map[int][]pick{}

	valid := []bool{}
	i := 0

	// p2
	total2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		game, results, _ := strings.Cut(line, ":")
		id, err := strconv.Atoi(strings.Split(game, " ")[1])
		if err != nil {
			panic(err)
		}
		valid = append(valid, true)
		p := []pick{}

		// p2 count the fewer color necessary to play each game
		red, green, blue := 0, 0, 0

		// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		for results != "" {
			var result string
			g := pick{}
			result, results, _ = strings.Cut(results, ";")
			// 3 blue, 4 red
			for _, picks := range strings.Split(result, ",") {
				pick := strings.Split(picks, " ")
				color := pick[2]
				count, err := strconv.Atoi(pick[1])
				if err != nil {
					panic(err)
				}
				switch color {
				case "red":
					g.red = count
					// p2
					if count > red {
						red = count
					}
				case "green":
					g.green = count
					// p2
					if count > green {
						green = count
					}
				case "blue":
					g.blue = count
					// p2
					if count > blue {
						blue = count
					}
				}
			}
			// p1
			p = append(p, g)
			if g.red > maxRed || g.green > maxGreen || g.blue > maxBlue {
				valid[i] = false
			}

		}
		// p1
		picks[id] = p
		if valid[i] {
			total += id
			fmt.Printf("valid play %d =  %v\n", id, p)
		}
		i += 1

		// p2
		fmt.Printf("Game %d: red = %d, green %d, blue %d\n", id, red, green, blue)
		total2 += red * green * blue
	}
	fmt.Println(total, total2)
}
