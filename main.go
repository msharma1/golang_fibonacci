package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

func generateFibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("invalid input, please enter a non-negative integer: %d", n)
	}
	if n == 0 {
		return []int{}, nil
	}
	if n == 1 {
		return []int{0}, nil
	}
	fibonacciArray := make([]int, n)
	fibonacciArray[0] = 0
	fibonacciArray[1] = 1
	for i := 2; i < n; i++ {
		a := fibonacciArray[i-1]
		b := fibonacciArray[i-2]

		if a > 0 && b > 0 && a > math.MaxInt-b {
			return nil, fmt.Errorf("integer overflow detected at term %d", i+1)
		}

		fibonacciArray[i] = a + b
	}
	return fibonacciArray, nil
}

func main() {
	log.SetFlags(0)
	var n int
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Generates the Fibonacci sequence up to a specified number of terms.")
		fmt.Println("\nOptions:")

		flag.PrintDefaults()
	}
	flag.IntVar(&n, "n", 0, "The number of Fibonacci results to return.")
	flag.IntVar(&n, "results", 0, "The number of Fibonacci results to return.")
	flag.Parse()
	if n <= 0 {
		flag.Usage()
		log.Fatal("\nError: Please provide a positive number using the -n flag.")
	}
	fibonacciNumbers, err := generateFibonacci(n)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("Fibonacci numbers:", fibonacciNumbers)
}
