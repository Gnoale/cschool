package combinatoire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var smallNumTestCases = []struct {
	pattern string
	want    string
}{
	{pattern: "IIIDIDDD", want: "123549876"},
	{pattern: "DDD", want: "4321"},
	{pattern: "I", want: "12"},
}

var permutationTestCases = []struct {
	numbers string
	k       int
	want    int
}{
	{numbers: "123456789", k: 9, want: 362880},
	{numbers: "123456789", k: 4, want: 3024},
}

func TestPermute(t *testing.T) {
	for _, tc := range permutationTestCases {
		results := []string{}
		UniquePermute(tc.numbers, "", &results, make(map[rune]bool), tc.k)
		assert.Equal(t, tc.want, len(results))
	}
}

func TestSmallestNumber(t *testing.T) {
	for _, tc := range smallNumTestCases {
		assert.Equal(t, tc.want, smallestNumber(tc.pattern))
	}
}
