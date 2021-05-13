package main

import (
	"fmt"

	"github.com/mdegaris/go-learning/factorial"
	"github.com/mdegaris/go-learning/greeting"
	"github.com/mdegaris/go-learning/primes"
)

func main() {
	greeting.Greet(greeting.ENGLISH)
	greeting.Greet(greeting.FRENCH)
	greeting.Greet(greeting.SPANISH)

	p := 300
	f := 7
	fmt.Println("Primes list", p, primes.GeneratePrimes(p))
	fmt.Println("Factorial of", f, factorial.Factorial(f))

}
