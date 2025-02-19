package combinatoire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var happyStringTestCases = []struct {
	n    int
	k    int
	want string
}{
	{n: 1, k: 3, want: "c"},
	{n: 1, k: 4, want: ""},
	{n: 3, k: 9, want: "cab"},
	// {n: 2, k: 7, want: ""},
	// {n: 10, k: 100, want: "abacbabacb"},
}

func TestHappyString(t *testing.T) {
	for _, tc := range happyStringTestCases {
		assert.Equal(t, tc.want, getHappyString(tc.n, tc.k))
	}
}
