package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasLeadingZero(t *testing.T) {
	s1 := "0000abc"
	r1 := HasLeadingZero(s1, 4)
	assert.Equal(t, true, r1, "0000abc has 4 leading zero")

	s2 := "0a0aa"
	r2 := HasLeadingZero(s2, 2)
	assert.Equal(t, false, r2, "0a0aa doesn't have 2 leading zero")
}
