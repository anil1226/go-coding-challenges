package goroutines

import (
	"math"
)

func isPrime(num int, ch chan bool) {
	if num < 2 {
		ch <- false
		return
	}

	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			ch <- false
			return
		}
	}

	ch <- true
}

func IsPrimeGo(num int) bool {
	ch := make(chan bool)

	go isPrime(num, ch)

	return <-ch
}
