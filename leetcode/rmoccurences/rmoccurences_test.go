package rmoccurences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	s        string
	part     string
	expected string
}{
	{
		s:        "daabcbaabcbc",
		part:     "abc",
		expected: "dab",
	},
	{
		s:        "axxxxyyyyb",
		part:     "xy",
		expected: "ab",
	},
	{
		s:        "ccctltctlltlb",
		part:     "ctl",
		expected: "b",
	},
	{
		s:        "eemckxmckx",
		part:     "emckx",
		expected: "",
	},
}

func TestRemoveOccurrences(t *testing.T) {
	for _, testCase := range testCases {
		result := RemoveOccurrences(testCase.s, testCase.part)
		assert.Equal(t, testCase.expected, result)
	}
}
