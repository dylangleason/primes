package primes

// Number is a simple wrapper for an integer type that also indicates
// whether the given number is prime or not, which can be useful for
// display purposes.
type Number struct {
	Number  int
	IsPrime bool
}

// NPrimes impelments an incremental sieve. See the following:
// https://www.cs.hmc.edu/~oneill/papers/Sieve-JFP.pdf
func NPrimes(n int) []Number {
	nums := []Number{}

	gen := generatePrime()
	for i := 0; i < n; i++ {
		nums = append(nums, Number{Number: <-gen, IsPrime: true})
	}

	return nums
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
	for i := 0; i <= n; i++ {
		primes[i] = Number{Number: i, IsPrime: true}
	}

	// Starting at prime number 2, for each p, remove all
	// composites. Stop processing when p^2 exceeds N.
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

// generatePrime uses channels to lazily compute an unbounded number
// of primes, returning a send channel of type int
func generatePrime() <-chan int {
	const firstPrime = 2

	result := make(chan int)
	composites := make(map[int][]int)

	go func() {
		defer close(result)

		for candidate := firstPrime; ; candidate++ {
			// Check to see if the current candidate is in
			// the composites table. If not, the candidate
			// is a prime number.
			factors, found := composites[candidate]
			if !found {
				result <- candidate
				// Store the prime number as a factor for the first
				// known multiple in a new array, which is the
				// product of the prime multiplied by itself.
				composites[candidate*candidate] = []int{candidate}
				continue
			}

			// Otherwise, the candidate is a known
			// composite or multiple found in the
			// composites table. Now consider all of the
			// factors seen thus far for the composite to
			// determine additional multiples to mark as
			// composite.
			for _, factor := range factors {
				// For each factor, the next multiple to mark as composite
				// is the factor plus the candidate itself.
				next := candidate + factor

				// If the composite already exists, append the
				// new factor, otherwise create a new array for this
				// composite including the factor.
				if _, found := composites[next]; found {
					composites[next] = append(composites[next], factor)
				} else {
					composites[next] = []int{factor}
				}
			}

			// Once done, remove the candidate to reduce
			// space complexity since we only need to keep
			// track of multiples for the next candidate
			// (i.e. lazily)
			delete(composites, candidate)
		}
	}()

	return result
}

// calcComposites will compute a sequence of composite
// (i.e. non-prime) numbers in range (n, max]. A composite is each
// successive increment of the given number, not including the number
// itself.
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
