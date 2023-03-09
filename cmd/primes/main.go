package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"

	"github.com/dylangleason/primes"
)

const tableWidth = 10

func main() {
	printTitle()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pick a number: ")
		num, err := readPrime(r)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		fmt.Printf("Displaying numbers 1-%d:\n\n", num)
		printPrimes(primes.PrimesUpTo(num))
	}
}

func readPrime(reader *bufio.Reader) (int, error) {
	in, _, err := reader.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(in))
}

func printTitle() {
	c := color.New(color.FgHiWhite).Add(color.Underline).Add(color.Bold)
	c.Printf("\nPrime Number Generator\n\n")
}

func printPrimes(p []primes.Number) {
	var (
		green = color.New(color.FgGreen)
		red   = color.New(color.FgRed)
	)

	for i, num := range p {
		if i > 0 && i%tableWidth == 0 {
			fmt.Println()
		}
		c := red
		if num.IsPrime {
			c = green
		}
		c.Printf("%d\t", num.Number)
	}

	fmt.Printf("\n\n")
}
