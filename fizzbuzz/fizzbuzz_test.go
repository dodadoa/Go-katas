package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsNotFizzbuzz(t *testing.T) {
	assert.Equal(t, "1", FizzBuzz(1))
  assert.Equal(t, "2", FizzBuzz(2))
  assert.Equal(t, "4", FizzBuzz(4))
}

func Test_IsFizz(t *testing.T) {
  assert.Equal(t, "Fizz", FizzBuzz(3))
  assert.Equal(t, "Fizz", FizzBuzz(6))
  assert.Equal(t, "Fizz", FizzBuzz(9))
	assert.Equal(t, "Fizz", FizzBuzz(31))
	assert.Equal(t, "Fizz", FizzBuzz(32))
}

func Test_IsBuzz(t *testing.T) {
  assert.Equal(t, "Buzz", FizzBuzz(5))
	assert.Equal(t, "Buzz", FizzBuzz(20))
	assert.Equal(t, "Buzz", FizzBuzz(25))
}

func Test_IsFizzBuzz(t *testing.T) {
  assert.Equal(t, "FizzBuzz", FizzBuzz(15))
	assert.Equal(t, "FizzBuzz", FizzBuzz(30))
	assert.Equal(t, "FizzBuzz", FizzBuzz(35))
	assert.Equal(t, "FizzBuzz", FizzBuzz(51))
	assert.Equal(t, "FizzBuzz", FizzBuzz(53))

}
