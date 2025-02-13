package cleardigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClearDigits(t *testing.T) {
	s := "cb34"
	s = clearDigits(s)
	assert.Equal(t, "", s)

	s = "cbdef"
	s = clearDigits(s)
	assert.Equal(t, "cbdef", s)

	s = "cbd4f"
	s = clearDigits(s)
	assert.Equal(t, "cbf", s)
}
