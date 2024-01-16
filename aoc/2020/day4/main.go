package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Gnoale/adventofcode/puzzlein"
)

var mFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func checkField(l []string) bool {
	// [hgt:160cm cid:103 ecl:gry byr:2000 hcl:#a97842 eyr:2026 pid:528503931 iyr:2010]
	m := make(map[string]string)
	for _, f := range l {
		kv := strings.Split(f, ":")
		m[kv[0]] = kv[1]
	}
	for _, f := range mFields {
		if d, found := m[f]; found {
			switch f {
			case "byr":
				di, _ := strconv.Atoi(d)
				if di < 1920 || di > 2002 {
					return false
				}
			case "iyr":
				di, _ := strconv.Atoi(d)
				if di < 2010 || di > 2020 {
					return false
				}
			case "eyr":
				di, _ := strconv.Atoi(d)
				if di < 2020 || di > 2030 {
					return false
				}
			case "hgt":
				recm, _ := regexp.Compile(`[\d]{3}?cm`)
				rein, _ := regexp.Compile(`[\d]{2}?in`)
				if recm.MatchString(d) {
					di, _ := strconv.Atoi(d[:len(d)-2])
					if di < 150 || di > 193 {
						return false
					}
				} else if rein.MatchString(d) {
					di, _ := strconv.Atoi(d[:len(d)-2])
					if di < 59 || di > 76 {
						return false
					}
				} else {
					return false
				}
			case "hcl":
				re, _ := regexp.Compile(`\#[0-9A-Fa-f]{6}?`)
				if !re.MatchString(d) {
					return false
				}
			case "ecl":
				if d != "amb" && d != "blu" && d != "brn" && d != "gry" && d != "grn" && d != "hzl" && d != "oth" {
					return false
				}
			case "pid":
				re, _ := regexp.Compile(`[0-9]{9}?`)
				if !re.MatchString(d) {
					return false
				}
			}
		} else {
			return false
		}
	}
	fmt.Println(m)
	return true
}

func checkValid(passp [][]string) int {
	var n int
	for _, pass := range passp {
		if checkField(pass) {
			n++
		}
	}
	return n
}

func main() {
	passp, err := puzzlein.GetStrEl("./day4/input")
	if err != nil {
		panic(err)
	}

	cpassp := puzzlein.Clean(passp)
	// part1
	n := checkValid(cpassp)
	fmt.Printf("total valid = %v\n", n)
}
