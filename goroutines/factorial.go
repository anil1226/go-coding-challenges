package goroutines

//Factorial Calculation with goroutines
func FactorialGo(num int) int {
	if num <= 0 {
		return 0
	}
	factorial := 1
	ch := make(chan int)
	for i := num; i > 0; i-- {
		go multiply(ch, factorial, i)
		factorial = <-ch
	}
	return factorial
}

func multiply(ch chan int, fact int, num int) {
	ch <- fact * num
}
