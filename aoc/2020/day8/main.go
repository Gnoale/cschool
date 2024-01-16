package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gnoale/adventofcode/puzzlein"
)

type intcode struct {
	input [][]string
	acc   int
	p     int
	size  int
	seen  []bool
}

func NewIntcode(in []string) *intcode {
	code := make([][]string, 0)
	for _, l := range in {
		code = append(code, strings.Split(l, " "))
	}
	seen := make([]bool, len(code))
	return &intcode{code, 0, 0, len(code), seen}
}

func (bc *intcode) check(i int) error {
	if bc.p+i >= bc.size {
		return nil
	}
	if bc.seen[bc.p+i] == true {
		return fmt.Errorf("error %v", strings.Join(bc.input[bc.p], " "))
	}
	return nil
}

func (bc *intcode) reset() {
	bc.p = 0
	bc.seen = make([]bool, bc.size)
	bc.acc = 0
}

func (bc *intcode) exec(cmd string, args string) error {
	sign := string(args[0])
	n, _ := strconv.Atoi(string(args[1:]))
	switch cmd {
	case "acc":
		if sign == "+" {
			bc.acc += n
		} else {
			bc.acc -= n
		}
		if err := bc.check(1); err != nil {
			return err
		}
		bc.p++
	case "jmp":
		var err error
		if sign == "+" {
			err = bc.check(n)
			bc.p += n
		} else {
			err = bc.check(-n)
			bc.p -= n
		}
		if err != nil {
			return err
		}
	default:
		bc.p++
	}
	fmt.Printf("%v %v p:%v acc:%v\n", cmd, args, bc.p, bc.acc)
	return nil
}

func (bc *intcode) run() error {
	var err error
	for bc.p < bc.size {
		bc.seen[bc.p] = true
		cmd := bc.input[bc.p][0]
		args := bc.input[bc.p][1]
		if err = bc.exec(cmd, args); err != nil {
			break
		}
	}
	return err

}

// part2
func (bc *intcode) bruteForce() {

	for i := 0; i < bc.size; i++ {
		var run bool
		var err error
		replaced := make(map[string]int)
		if bc.input[i][0] == "jmp" {
			bc.input[i][0] = "nop"
			replaced["jmp"] = i
			run = true
		} else if bc.input[i][0] == "nop" {
			bc.input[i][0] = "jmp"
			replaced["nop"] = i
			run = true
		}
		if run {
			err = bc.run()
			if err != nil {
				for k, v := range replaced {
					bc.input[v][0] = k
				}
				bc.reset()
			} else {
				panic("win")
			}
		}
	}
}

func main() {
	in, err := puzzlein.GetStr("./day8/input")
	if err != nil {
		panic(err)
	}

	bootcode := NewIntcode(in)
	//bootcode.run()
	bootcode.bruteForce()

}
