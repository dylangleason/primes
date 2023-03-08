package nprimes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNPrimes(t *testing.T) {
	tests := map[string]struct {
		input  int
		output []int
	}{
		"10 primes": {
			input:  10,
			output: []int{2, 3, 5, 7},
		},
		"25 primes": {
			input:  25,
			output: []int{2, 3, 5, 7, 11, 13, 17, 19, 23},
		},
		"50 primes": {
			input:  50,
			output: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}

	assert := assert.New(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.EqualValues(test.output, NPrimes(test.input))
		})
	}
}
