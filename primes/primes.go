package primes

import (
	"strconv"
)

func GeneratePrimes(n int) []string {

	ch := make(chan int)
	go nextPrime(ch)

	primes := []string{}
	for i := 0; i < n; i++ {
		prime := strconv.Itoa(<-ch)
		primes = append(primes, prime)
	}
	return primes
}

func nextPrime(ch chan int) {
	for i := 2; ; i++ {
		isPrime := true
		for j := 1; j < i; j++ {
			if (j > 1) && ((i % j) == 0) {
				isPrime = false
			}
			if !isPrime {
				break
			}
		}
		if isPrime {
			ch <- i
		}
	}
}
