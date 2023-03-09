package primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNPrimes(t *testing.T) {
	tests := map[string]struct {
		input  int
		output []int
	}{
		"5 primes": {
			input:  5,
			output: []int{2, 3, 5, 7, 11},
		},
		"10 primes": {
			input:  10,
			output: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		"15 primes": {
			input:  15,
			output: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}

	assert := assert.New(t)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := NPrimes(test.input)
			assert.EqualValues(test.output, onlyPrimesAsInts(result))
		})
	}

}

func TestPrimesUpTo(t *testing.T) {
	tests := map[string]struct {
		input  int
		output []int
	}{
		"Up to 10": {
			input:  10,
			output: []int{2, 3, 5, 7},
		},
		"Up to 25": {
			input:  25,
			output: []int{2, 3, 5, 7, 11, 13, 17, 19, 23},
		},
		"Up to 50": {
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

const excessivelyLargeNPrimes = 125_000

var primesUpToResult []Number

func BenchmarkPrimesUpTo(b *testing.B) {
	var r []Number
	for n := 0; n < b.N; n++ {
		r = PrimesUpTo(excessivelyLargeNPrimes)
	}
	primesUpToResult = r
}

var nPrimesResult []Number

func BenchmarkNPrimes(b *testing.B) {
	var r []Number
	for n := 0; n < b.N; n++ {
		r = NPrimes(excessivelyLargeNPrimes)
	}
	nPrimesResult = r
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
