package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(JoinInt([]int{1, 2, 3}, " "), "1 2 3")
	assert.NotEqual(JoinInt([]int{1, 2, 3}, " "), "1 2 3 ")

	assert.True(MemberOfSlice(1, []int{1, 2, 3}))
	assert.False(MemberOfSlice(1, []int{2, 3, 4}))
	assert.Panics(func() { MemberOfSlice('a', []int{1, 2, 3}) },
		"First argument and slice elements have different types")
	assert.Panics(func() { MemberOfSlice(-1, []uint{1, 2, 3}) },
		"First argument and slice elements have different types")
	assert.Panics(func() { MemberOfSlice(1, 2) },
		"Second argument is not a slice")
}
