package combinatoire

import (
	"fmt"
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

var uniquePermuteTestCases = []struct {
	numbers string
	k       int
	want    int
}{
	{numbers: "123456789", k: 9, want: 362880},
	{numbers: "123456789", k: 4, want: 3024},
}

var permuteTestCases = []struct {
	numbers string
	k       int
	want    int
}{
	{numbers: "10", k: 2, want: 4}, // 10, 01, 00 ,11
	{numbers: "01", k: 4, want: 16},
}

func TestPermute(t *testing.T) {
	for _, tc := range uniquePermuteTestCases {
		results := []string{}
		UniquePermute(tc.numbers, "", &results, make(map[rune]bool), tc.k)
		assert.Equal(t, tc.want, len(results))
	}

	for _, tc := range permuteTestCases {
		results := []string{}
		Permute(tc.numbers, "", &results, tc.k)
		assert.Equal(t, tc.want, len(results))
		fmt.Println(results)
	}
}

func TestSmallestNumber(t *testing.T) {
	for _, tc := range smallNumTestCases {
		assert.Equal(t, tc.want, smallestNumber(tc.pattern))
	}
}

func TestFindDifferentBinaryString(t *testing.T) {
	nums := []string{"01", "10"}
	assert.Equal(t, "00", findDifferentBinaryString(nums))
	nums = []string{"011", "001", "000"}
	assert.Equal(t, "010", findDifferentBinaryString(nums))
}
