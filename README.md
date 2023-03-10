# Prime Numbers

Routines for generating prime numbers.

## Installation

Assuming you have Docker, you can run the below `make` targets to run the program in a Docker container.

## Running the Program

You can use the Dockerfile to build and execute the program thusly:

```console
make
```

Make sure to use a terminal emulator that supports [ANSI Colors](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors), as the printed primes are indicated in green color.

## Running The Tests

You can run the unit tests using the following:

```console
make test
```

Benchmark tests can also be run for the `PrimesUpTo` function:

```console
make bench
```

## Learning

- I started off by trying to understand a bounded prime number generator using a the classic Sieve of Eratosthenes, via implementation `PrimesUpTo`, which just computes primes up to a number (inclusive).
- Given that the multiplication table is unbounded, I tried to implement an unbounded sieve using an arbitrary scaling factor via `NPrimes` for computing each composite, but found this wouldn't work and instead needed to investigate an "incremental sieve" that used infinite sequences.
- I realized that this problem would have been easier to solve using a language that has the concept of a lazy or infinite sequence, which Go does not support out of the box.

## Pros

- I think the classic sieve solution is well commented and easy to reason about.
- I used a test driven development approach for prime number generation to validate each step and test my understanding.
- Use of `Number` wrapper struct provides some metadata that allowed me to decouple presentation from logic in the case of displaing `PrimesUpTo`, though that isn't relevant to the original problem description and probably unnecessary.

## Cons

- I had to generate the multiplication table output using my sub-standard `NPrimes` implementation, which isn't a true incremental sieve
- Go does not support the notion of a lazy / infinite sequence so I got stuck trying to find a solution using primitives for the latter
- Console output for table probably should have been built using a StringBuilder so I could test multiplication table output via unit tests
