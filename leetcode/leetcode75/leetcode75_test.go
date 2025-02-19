package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
	}

	for _, tt := range tests {
		got := maxArea(tt.height)
		assert.Equal(t, tt.want, got, "maxArea(%v) = %v, want %v", tt.height, got, tt.want)
	}
}

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 0, 1}, 3},
		{[]int{1, 1, 1}, 2},
	}

	for _, tt := range tests {
		got := longestSubarray(tt.nums)
		assert.Equal(t, tt.want, got, "longestSubarray(%v) = %v, want %v", tt.nums, got, tt.want)
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, tt := range tests {
		got := reverseWords(tt.input)
		assert.Equal(t, tt.want, got, "reverseWords(%q) = %q, want %q", tt.input, got, tt.want)
	}
}

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{1, 2, 3, 4}, 2, 0},
	}

	for _, tt := range tests {
		got := maxOperations(tt.nums, tt.k)
		assert.Equal(t, tt.want, got, "maxOperations(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		got := productExceptSelf(tt.nums)
		assert.Equal(t, tt.want, got, "productExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
	}
}

func TestCompress(t *testing.T) {
	tests := []struct {
		input []byte
		want  int
	}{
		{[]byte("aabbccc"), 6},
		{[]byte("a"), 1},
		{[]byte("abbbbbbbbbbbbb"), 4},
	}

	for _, tt := range tests {
		got := compress(tt.input)
		assert.Equal(t, tt.want, got, "compress(%s) = %d, want %d", tt.input, got, tt.want)
	}
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		want         []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	for _, tt := range tests {
		got := dailyTemperatures(tt.temperatures)
		assert.Equal(t, tt.want, got, "dailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
	}
}
