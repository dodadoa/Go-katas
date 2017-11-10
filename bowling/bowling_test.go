package bowling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Frame(t *testing.T) {
	assert.Equal(t, []int{3, 8}, frame([]int{1, 2, 3, 5}))
	assert.Equal(t, []int{15, 5, 7}, frame([]int{10, 2, 3, 7}))
	assert.Equal(t, []int{12, 2}, frame([]int{10, 2}))
}

func Test_Score(t *testing.T) {
	assert.Equal(t, 300, score([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}))
	assert.Equal(t, 90, score([]int{9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0, 9, 0}))
	assert.Equal(t, 150, score([]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}))
}
