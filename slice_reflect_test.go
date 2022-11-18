//go:build !go1.18 || reflect

package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceReflect(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() { MemberOfSlice('a', []int{1, 2, 3}) },
		"First argument and slice elements have different types")
	assert.Panics(func() { MemberOfSlice(-1, []uint{1, 2, 3}) },
		"First argument and slice elements have different types")
	assert.Panics(func() { MemberOfSlice(1, 2) },
		"Second argument is not a slice")
}
