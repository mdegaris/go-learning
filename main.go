package main

import (
	"fmt"

	"github.com/mdegaris/go-learning/factorial"
	"github.com/mdegaris/go-learning/greeting"
	"github.com/mdegaris/go-learning/primes"
)

func main() {
	greeting.English()
	greeting.French()

	fmt.Println(primes.GeneratePrimes(10))
	fmt.Println(factorial.Factorial(5))

}
