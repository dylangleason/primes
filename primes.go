package primes

// Number is a simple wrapper for an integer type that also indicates
// whether the given number is prime or not, which can be useful for
// display purposes.
type Number struct {
	Number     int
	IsPrime    bool
	Composites []int
}

// NPrimes should use an incremental sieve. I tried to follow this but
// couldn't quite figure out what was being done:
// https://www.cs.hmc.edu/~oneill/papers/Sieve-JFP.pdf
func NPrimes(n int) []Number {
	primes := make([]Number, 0, n)

	// Use a set to keep track of numbers containing
	// composite numbers, aka not prime.
	composites := make(map[int]struct{})

	number := 2
	count := 0

	for len(primes) < n {
		// This maximum condition isn't correct, just faking
		// it here. While this will produce functionally
		// correct results, it's much too slow for large
		// primes. The benchmark test will time out.
		max := n * n

		// TODO: hold each composite for the given number so
		// they can be used in the table later.

		calcComposites(number, max, func(composite int) {
			composites[composite] = struct{}{}
			count++
		})

		if _, found := composites[number]; !found {
			primes = append(primes, Number{Number: number, IsPrime: true})
		}

		count = 0
		number++
	}

	return primes
}

// PrimesUpTo will generate an array of integers up to a maximum
// number N using the "Sieve of Eratosthenes" algorithm for efficient
// prime number generation.
func PrimesUpTo(n int) []Number {
	// Create a range of boolean values indexed from 2 to N+1, in
	// order to keep track of those that are prime. Note that this
	// makes indexing 2-N easier when calculating primes, but it
	// does mean that index 0 (p=1) will go unused.
	primes := make([]Number, n+1)
	for i := 0; i < n; i++ {
		primes[i] = Number{Number: i, IsPrime: true}
	}

	// Starting at prime number 2, for each p, remove all
	// non-primes. Stop processing when p^2 exceeds N.
	for p := 2; p*p <= n; p++ {
		if primes[p].IsPrime {
			calcComposites(p, n, func(composite int) {
				primes[composite].IsPrime = false
			})
		}
	}

	// zero out first two elements, as they are not needed, before
	// reslicing, as both elements to allow garbage collection
	primes[0], primes[1] = Number{}, Number{}
	return primes[2:]
}

// calcComposites will compute a sequence of non-prime numbers in
// range (n, max]. A non-prime is each successive increment of the
// given number, not including the number itself.
//
// Example: num = 2, max = 10
//
// nonPrime = 2 * 2 = 4      NOTE: initialize to n^2 because we don't count n itself
// primes[nonPrime] = false
//
// nonPrime = 4 + 2 = 6
// primes[6] = false
//
// nonPrime = 6 + 2 = 8
// primes[nonPrime] = false
//
// nonPrime = 8 + 2 = 10
// primes[nonPrime] = false
func calcComposites(num, max int, callback func(nonPrime int)) {
	for nonPrime := num * num; nonPrime <= max; nonPrime += num {
		callback(nonPrime)
	}
}
