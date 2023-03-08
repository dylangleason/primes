package nprimes

// NPrimes will generate N primes as an array of integers using the
// "Sieve of Eratosthenes" algorithm for efficient prime number
// generation.
func NPrimes(n int) []int {
	// Create a range of boolean values indexed from 2 to N+1, in
	// order to keep track of those that are prime. Note that this
	// makes indexing 2-N easier when calculating primes, but it
	// does mean that index 0 (p=1) will go unused.
	primes := make([]bool, n+1)
	for i := 0; i < n; i++ {
		primes[i] = true
	}

	// Starting at prime number 2, for each p, remove all
	// non-primes. Stop processing when p^2 exceeds N.
	for p := 2; p*p <= n; p++ {
		// TODO(dylangleason): implement non-prime removal logic
	}

	// Add each valid prime to the numbers array. Note that we
	// need to reallocate since the number of primes is
	// non-deterministic. We could possibly set an initial
	// capacity based on ratio of primes to N (logarithmic
	// output?) to reduce space complexity of final result without
	// repeated allocations.
	nums := []int{}
	for i := 2; i <= n; i++ {
		if primes[i] {
			nums = append(nums, i)
		}
	}

	return nums
}
