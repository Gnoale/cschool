package combinatoire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	digits string
	want   []string
}{
	{digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{digits: "234", want: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
	{digits: "2", want: []string{"a", "b", "c"}},
	{digits: "", want: []string{}},
}

func TestLetterCombinations(t *testing.T) {
	for _, tc := range testCases {
		assert.Equal(t, tc.want, letterCombinations(tc.digits))
	}
}
