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
