package nprimes

// Number is a simple wrapper for an integer type that also indicates
// whether the given number is prime or not, which can be useful for
// display purposes.
type Number struct {
	Number  int
	IsPrime bool
}

// NPrimes will generate N primes as an array of integers using the
// "Sieve of Eratosthenes" algorithm for efficient prime number
// generation.
func NPrimes(n int) []Number {
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
			removeNonPrime(p, n, primes)
		}
	}

	// zero out first two elements, as they are not needed, before
	// reslicing, as both elements to allow garbage collection
	primes[0], primes[1] = Number{}, Number{}
	return primes[2:]
}

// removeNonPrime will set all indexes corresponding to numbers in
// range (n, max] to false, indicating those numbers as non-prime. A
// non-prime is each successive increment of the given number, not
// including the number itself.
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
func removeNonPrime(num, max int, primes []Number) {
	for nonPrime := num * num; nonPrime <= max; nonPrime += num {
		primes[nonPrime].IsPrime = false
	}
}
