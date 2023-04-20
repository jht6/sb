package block

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	blk := New(
		0,
		"a",
		[]string{"b"},
		-1,
	)
	hash := blk.Hash()
	assert.NotEqual(t, "", hash)
}
