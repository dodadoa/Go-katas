package bowling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_try(t *testing.T) {
	assert.Equal(t, 10, score([4]int{1, 2, 3, 4}))
}
