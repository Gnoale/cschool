package aoc

import (
	"2023/pkg/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type race struct {
	time     int
	distance int
}

//func parse(scanner *bufio.Scanner) []race {
//	races := []race{}
//	i := 0
//	for scanner.Scan() {
//		//	races = append(races, race)
//		line := scanner.Text()
//		// races time
//		if i == 0 {
//			sline := strings.Split(line, " ")
//			for i := 1; i < len(sline); i++ {
//				if sline[i] == "" {
//					continue
//				}
//				t, err := strconv.Atoi(strings.Trim(sline[i], " "))
//				if err != nil {
//					panic(err)
//				}
//				races = append(races, race{
//					time: t,
//				})
//
//			}
//		}
//		// races distance
//		if i == 1 {
//			sline := strings.Split(line, " ")
//			j := 0
//			for i := 1; i < len(sline); i++ {
//				if sline[i] == "" {
//					continue
//				}
//				t, err := strconv.Atoi(strings.Trim(sline[i], " "))
//				if err != nil {
//					panic(err)
//				}
//				races[j].distance = t
//				j++
//			}
//		}
//		i++
//	}
//	return races
//}

func parseDay6P2(scanner *bufio.Scanner) race {
	r := race{}
	scanner.Scan()
	line := scanner.Text()
	// race time
	sline := strings.Split(line, ":")[1]
	num := 0
	j := 0
	for len(sline) > 0 {
		r, size := utf8.DecodeLastRuneInString(sline)
		if size == 0 {
			sline = sline[:len(sline)-1]
			continue
		}
		t, err := strconv.Atoi(string(r))
		if err != nil {
			sline = sline[:len(sline)-size]
			continue
		}
		num += t * utils.Pow(10, j)
		j++
		sline = sline[:len(sline)-size]
	}
	r.time = num

	scanner.Scan()
	line = scanner.Text()
	// races distance

	sline = strings.Split(line, ":")[1]
	num = 0
	j = 0
	for len(sline) > 0 {
		r, size := utf8.DecodeLastRuneInString(sline)
		if size == 0 {
			sline = sline[:len(sline)-1]
			continue
		}
		t, err := strconv.Atoi(string(r))
		if err != nil {
			sline = sline[:len(sline)-size]
			continue
		}
		num += t * utils.Pow(10, j)
		j++
		sline = sline[:len(sline)-size]
	}
	r.distance = num
	return r
}

func Race(input *bufio.Scanner) {
	race := parseDay6P2(input)
	fmt.Println(race)
	/*
		For each race, find the different amount of time you need to push the button to win the race
		let k = time we hold the button
		let d the distance
		let n the total time
		d = k * n-k

	*/
	total := 0
	for i := 1; i < race.time; i++ {

		d := i * (race.time - i)
		if d > race.distance {
			total += 1
		}
	}
	fmt.Println(total)

}
