package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz(t *testing.T) {
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"}
	chain, err := Fizzbuzz("fizz", "buzz", 3, 5, 10)
	assert.NoError(t, err)
	assert.Equal(t, expected, chain)
}

func TestFizzbuzzParams(t *testing.T) {
	_, err := Fizzbuzz("fizz", "buzz", 3, 5, -10)
	assert.Error(t, err)

	_, err = Fizzbuzz("fizz", "buzz", 3, -29, 10)
	assert.Error(t, err)

	chain, err := Fizzbuzz("fizz", "buzz", 14, 37, 99999)
	assert.NoError(t, err)
	assert.NotNil(t, chain)
}
