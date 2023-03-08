package primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimesUpTo(t *testing.T) {
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
			result := PrimesUpTo(test.input)
			assert.EqualValues(test.output, onlyPrimesAsInts(result))
		})
	}
}

var benchResult []Number

func BenchmarkNPrimes(b *testing.B) {
	const excessivelyLargeNPrimes = 125_000
	var r []Number
	for n := 0; n < b.N; n++ {
		r = PrimesUpTo(excessivelyLargeNPrimes)
	}
	benchResult = r
}

func onlyPrimesAsInts(nums []Number) []int {
	primes := []int{}
	for _, num := range nums {
		if num.IsPrime {
			primes = append(primes, num.Number)
		}
	}
	return primes
}
