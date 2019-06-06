package tdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	assert.NotEqual(t, "", Find("inroot"))
	assert.Equal(t, "", Find("inroot-not-exist"))
	assert.NotEqual(t, "", Find("subdir/insubdir"))
	assert.Equal(t, "", Find("subdir/insubdir-not-exist"))
}
