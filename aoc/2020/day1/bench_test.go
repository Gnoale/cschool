package main

import (
	"testing"

	"github.com/Gnoale/adventofcode/puzzlein"
)

var codes []int

func init() {
	var err error
	if codes, err = puzzlein.GetInt("input"); err != nil {
		panic(err)
	}
}

func BenchmarkFpn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findProductN(codes)
	}
}

func BenchmarkFpr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findProductNr(codes)
	}
}

func BenchmarkFpNn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findProductNn(codes)
	}
}
