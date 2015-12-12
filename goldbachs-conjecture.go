package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var primes = make([]int, 0, 4000)
var cprimes = make([]bool, 33000, 33000)

func sqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func isPrime(n int) bool {
	for i := 2; i <= sqrt(n); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func findPrimes() {
	primes = append(primes, 2)
	cprimes[2] = true
	for i := 3; i < 32000; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
			cprimes[i] = true
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var T int
	fmt.Fscan(reader, &T)
	findPrimes()

	for t := 0; t < T; t++ {
		var N int
		fmt.Fscan(reader, &N)
		if N < 4 {
			fmt.Fprint(writer, N, " has 0 representation(s)\n\n")
			continue
		}

		count := 0
		for _, p := range primes {
			if p > (N / 2) {
				break
			}

			if cprimes[N-p] {
				count++
			}
		}

		fmt.Fprint(writer, N, " has ", count, " representation(s)\n")
		for _, p := range primes {
			if p > (N / 2) {
				break
			}

			if cprimes[N-p] {
				fmt.Fprint(writer, p, "+", N-p, "\n")
			}
		}
		fmt.Fprint(writer, "\n")
	}

	writer.Flush()
}
