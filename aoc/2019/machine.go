package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

/*
Intcode machine

	Read a list of integer separated by comas

	1,9,10,3,2,3,11,0,99,30,40,50

		1,9,10,3
		ADD integers at position 9 and 10, store results in 3
		1,9,10,70

		Step fw 4 positions, etc.
*/

// var sample = "1,9,10,3,2,3,11,0,99,30,40,50"
var sample = "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,5,19,23,2,10,23,27,1,27,5,31,2,9,31,35,1,35,5,39,2,6,39,43,1,43,5,47,2,47,10,51,2,51,6,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,1,71,5,75,1,13,75,79,1,6,79,83,2,83,13,87,1,87,6,91,1,10,91,95,1,95,9,99,2,99,13,103,1,103,6,107,2,107,6,111,1,111,2,115,1,115,13,0,99,2,0,14,0"

// Opcode

const (
	ADD = iota + 1
	MUL
	STORE
	DIAG
	JMPT
	JMPF
	LESS
	EQUAL
	OFFSET

	EXIT = 99
)

type machine struct {
	codes   []int
	current int
	pointer int
}

func Newmachine(input string) *machine {
	var s scanner.Scanner
	//f, err := os.Open(codes)
	r := strings.NewReader(input)
	s.Init(r)

	codes := []int{}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		code, err := strconv.Atoi(s.TokenText())
		if err != nil {
			continue
		}
		codes = append(codes, code)
	}
	return &machine{
		codes: codes,
	}
}

func (l *machine) read(pos int) int {
	var v int
	if pos >= len(l.codes)-1 {
		v = l.codes[pos%len(l.codes)]
	} else {
		v = l.codes[pos]
	}
	return v
}

func (l *machine) move(pos int) {
	if pos >= len(l.codes)-1 {
		l.pointer = pos % len(l.codes)
	} else {
		l.pointer += pos
	}
	l.current = l.codes[l.pointer]
}

func (l *machine) set(pos int, value int) {
	if pos >= len(l.codes)-1 {
		l.codes[pos%len(l.codes)] = value
	} else {
		l.codes[pos] = value
	}
}

func (l *machine) init() {
	l.current = l.codes[l.pointer]
}

func (l *machine) exec() {
	for l.current != EXIT {
		switch l.current {
		case ADD:
			l.Add()
		case MUL:
			l.Mul()
		case STORE:
			l.Store()
		case DIAG:
			l.Diag()
		}
	}
	fmt.Println("EXIT ", l.codes)
}

// Add sum(&n+1, &n+2) = &n+3
func (l *machine) Add() {
	i1 := l.read(l.pointer + 1)
	i2 := l.read(l.pointer + 2)
	s := l.read(i1) + l.read(i2)
	r := l.read(l.pointer + 3)
	l.set(r, s)
	l.move(4)
}

func (l *machine) Mul() {
	i1 := l.read(l.pointer + 1)
	i2 := l.read(l.pointer + 2)
	s := l.read(i1) * l.read(i2)
	r := l.read(l.pointer + 3)
	l.set(r, s)
	l.move(4)
}

func (m *machine) Store() {
	i1 := l.read(l.read(l.pointer + 1))
	r := l.read(l.pointer + 2)
	l.set(r, i1)
	l.Move(3)

}

func (m *machine) Diag() {
	fmt.Printf("DIAG CODE %d\n", l.read(l.read(l.pointer+1)))
	l.Move(2)
}

func main() {
	l := Newmachine(sample)
	l.init()
	l.exec()
}
