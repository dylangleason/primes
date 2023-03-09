package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dylangleason/primes"
	"github.com/fatih/color"
)

const maxChoices = 2

const (
	choiceUpTo = iota + 1
	choiceNPrimesTable
)

func main() {
	printTitle()
	r := bufio.NewReader(os.Stdin)

	for {
		printMenu()

		choice := readChoice(r, maxChoices)
		num := readNumber(r)

		switch choice {
		case choiceUpTo:
			printPrimesUpTo(num)
		case choiceNPrimesTable:
			printFirstNPrimesTable(num)
		}
	}
}

func readSelection(reader *bufio.Reader) (int, error) {
	in, _, err := reader.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(in))
}

func readChoice(reader *bufio.Reader, maxChoices int) int {
	for {
		fmt.Print("\nChoose: ")
		choice, err := readSelection(reader)
		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
		if choice < 1 || choice > maxChoices {
			fmt.Printf("Invalid choice. Choose a number between 1-%d", maxChoices)
			continue
		}
		return choice
	}
}

func readNumber(reader *bufio.Reader) int {
	for {
		fmt.Print("\nPick a number: ")
		number, err := readSelection(reader)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}
		return number
	}
}

func printTitle() {
	c := color.New(color.FgHiWhite).Add(color.Underline).Add(color.Bold)
	c.Printf("\nPrime Number Generator\n\n")
}

func printMenu() {
	fmt.Println("What would you like to do?")
	fmt.Println("1. Count Primes Up To a Number")
	fmt.Println("2. View Multiplication Table For First N primes")
}

func printPrimesUpTo(number int) {
	const tableWidth = 10

	p := primes.PrimesUpTo(number)
	green, red := getColors()

	fmt.Printf("Numbers 1-%d:\n\n", number)
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

func printFirstNPrimesTable(count int) {
	p := primes.NPrimes(count)
	green, red := getColors()

	fmt.Printf("Multiplication table for first %d primes:\n\n", count)
	fmt.Printf(" \t")
	for _, prime := range p {
		green.Printf("%d\t", prime.Number)
	}
	fmt.Println()

	// This could probably be made a more efficient O(n) using an
	// offset..
	for i := 0; i < len(p); i++ {
		col := p[i].Number
		green.Printf("%d\t", col)
		for j := 0; j < len(p); j++ {
			red.Printf("%d\t", col*p[j].Number)
		}
		fmt.Println()
	}
	fmt.Println()
}

func getColors() (prime, comp *color.Color) {
	prime = color.New(color.FgGreen)
	comp = color.New(color.FgRed)
	return
}
