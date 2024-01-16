package aoc

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Seed(scanner *bufio.Scanner) {

	/* Seeds almanac

	seeds to be planted

	seed => soil, fertilizer
	fertilizer => water
	water =>  light
	light => temp
	temp => humidity
	humidity => location


	seed-to-soil map:
	50 98 2

		50 dst start, 98 src start, 2 lenght
		seed 98 => soil 50
		seed 99 => soil 51

	52 50 48
		48 values, seed 53 => soil 55


	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15

	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4

	etc.

	I need to build a map of each kind
	Then lookup maps consecutively starting from every initial seed
	and find the lowest location number

	*/

	seeds, maps := buildMaps(scanner)
	//buildMaps(scanner)
	visitSeeds(seeds, maps)

}

// visit build if required the index to visit the next map
func visitSeeds(seeds []int, maps map[string]preMap) {

	// build a list of maps in the expected order
	mapList := []preMap{}

	mapList = append(mapList, maps["soil"])
	mapList = append(mapList, maps["fertilizer"])
	mapList = append(mapList, maps["water"])
	mapList = append(mapList, maps["light"])
	mapList = append(mapList, maps["temperature"])
	mapList = append(mapList, maps["humidity"])
	mapList = append(mapList, maps["location"])

	//fmt.Println(mapList)

	min := 1000000000

	//visited := map[int]int{}

	// PART2 now seeds comes in pairs, 1st is the initial, 2nd is the range
	for i := 0; i < len(seeds); i += 2 {
		seedi := seeds[i]
		seedj := seeds[i+1]

		for j := seedi; j < seedi+seedj; j++ {
			source, dest := j, j

			//	runtime.Breakpoint()
			//fmt.Println("seed", j)
			for _, m := range mapList {
				/*
					seed-to-soil map:
					50 98 2
					52 50 48
				*/
				for i := 0; i < len(m.src); i++ {
					// 79 < 98
					// cannot be in the interval
					if source < m.src[i] {
						dest = source
						continue
					}
					// previous number is in the interval
					/*  seed 79, dst 52, src 50, step 48
					79 - 50 <= step
					*/
					diff := source - m.src[i]
					if diff <= m.steps[i] {
						dest = m.dst[i] + diff
						source = dest
						break
					}
				}
				//fmt.Println(m.key, dest)
			}
			//visited[j] = dest
			//fmt.Println()
			if dest < min {
				min = dest
			}

		}

	}

	fmt.Println(min)

}

type preMap struct {
	key   string
	src   []int
	dst   []int
	steps []int
}

// parse the input, build the maps
func buildMaps(scanner *bufio.Scanner) (seeds []int, maps map[string]preMap) {

	maps = map[string]preMap{}
	seeds = []int{}

	// get the seeds
	scanner.Scan()
	s := strings.Trim(strings.Split(scanner.Text(), ":")[1], " ")
	seedline := strings.Split(s, " ")
	// part1
	for _, num := range seedline {
		i, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, i)
	}

	//for i := 0; i < len(seedline); i += 2 {
	//	// now seeds comes in pairs, 1st is the initial, 2nd is the range
	//	seedi, err := strconv.Atoi(seedline[i])
	//	if err != nil {
	//		panic(err)
	//	}
	//	seedj, err := strconv.Atoi(seedline[i+1])
	//	if err != nil {
	//		panic(err)
	//	}
	//	for j := seedi; j < seedi+seedj; j++ {
	//		seeds = append(seeds, j)
	//	}
	//	fmt.Println(seedi)

	//}

	var key string
	currentMapper := preMap{}
	newMap := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "seed-to-soil") {
			key = "soil"
			newMap = true
		}

		if strings.HasPrefix(line, "soil-to-fertilizer") {
			key = "fertilizer"
			newMap = true
		}

		if strings.HasPrefix(line, "fertilizer-to-water") {
			key = "water"
			newMap = true
		}

		if strings.HasPrefix(line, "water-to-light") {
			key = "light"
			newMap = true
		}

		if strings.HasPrefix(line, "light-to-temperature ") {
			key = "temperature"
			newMap = true
		}

		if strings.HasPrefix(line, "temperature-to-humidity") {
			key = "humidity"
			newMap = true
		}

		if strings.HasPrefix(line, "humidity-to-location") {
			key = "location"
			newMap = true
		}

		if newMap {
			if len(currentMapper.src) != 0 {
				maps[currentMapper.key] = currentMapper
			}
			currentMapper = preMap{
				key: key,
			}
			newMap = false
			continue
		}

		numbers := strings.Split(line, " ")

		dstStart, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		srcStart, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		steps, err := strconv.Atoi(numbers[2])
		if err != nil {
			panic(err)
		}
		currentMapper.src = append(currentMapper.src, srcStart)
		currentMapper.dst = append(currentMapper.dst, dstStart)
		currentMapper.steps = append(currentMapper.steps, steps)

	}

	maps[currentMapper.key] = currentMapper
	fmt.Println(maps)

	//	line := scanner.Text()

	return seeds, maps
}
