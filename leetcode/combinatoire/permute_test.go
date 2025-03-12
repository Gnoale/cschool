package combinatoire

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var uniquePermuteTestCases = []struct {
	elements string
	k        int
	want     int
}{
	{elements: "123456789", k: 9, want: 362880},
	{elements: "123456789", k: 4, want: 3024},
}

var permuteTestCases = []struct {
	elements string
	k        int
	want     int
}{
	{elements: "10", k: 2, want: 4}, // 10, 01, 00 ,11
	{elements: "01", k: 4, want: 16},
}

func TestPermute(t *testing.T) {
	for _, tc := range uniquePermuteTestCases {
		results := []string{}
		UniquePermute(tc.elements, "", &results, make(map[rune]bool), tc.k)
		assert.Equal(t, tc.want, len(results))
	}

	for _, tc := range permuteTestCases {
		results := []string{}
		Permute(tc.elements, "", &results, tc.k)
		assert.Equal(t, tc.want, len(results))
		fmt.Println(results)
	}
}

var smallNumTestCases = []struct {
	pattern string
	want    string
}{
	{pattern: "IIIDIDDD", want: "123549876"},
	{pattern: "DDD", want: "4321"},
	{pattern: "I", want: "12"},
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

var changeTestCases = []struct {
	amount int
	coins  []int
	want   int
}{
	//{amount: 4, coins: []int{1, 2, 3}, want: 4},
	//	{amount: 245, coins: []int{16, 30, 9, 17, 40, 13, 42, 5, 25, 49, 7, 23, 1, 44, 4, 11, 33, 12, 27, 2, 38, 24, 28, 32, 14, 50}, want: 64027917156},
	//	{amount: 250, coins: []int{41, 34, 46, 9, 37, 32, 42, 21, 7, 13, 1, 24, 3, 43, 2, 23, 8, 45, 19, 30, 29, 18, 35, 11}, want: 15685693751},
	{amount: 8, coins: []int{1, 2, 3, 4}, want: 15},
}

func TestMakeChange(t *testing.T) {
	for _, tc := range changeTestCases {
		memo := make(map[string]int)
		ways := makeChange(tc.amount, tc.coins, 0, memo)
		assert.Equal(t, tc.want, ways)
	}
}

func TestMakeChangeDfs(t *testing.T) {
	for _, tc := range changeTestCases {
		memo := make(map[string]int)
		ways := makeChangeDfs(tc.amount, tc.coins, 0, 0, memo)
		assert.Equal(t, tc.want, ways)
		fmt.Println(memo)
	}
}
