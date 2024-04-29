// Coding Challenge
package main

import (
	"fmt"
	"go-Practice/goroutines"
)

func main() {
	// list := []int{10, 56, 28, 45, 87, 91, 42, 78, 99}
	// fmt.Println(evenSum(list))
	//fmt.Println(stringReverse("lina"))
	//fmt.Println(isPalindrome("madam"))
	//fibonacci(10)
	//fmt.Println(factorial(5))
	//fmt.Println(isPrime(14))
	fmt.Println(goroutines.FactorialGo(5))
	//app.MatrixProc()
}

func evenSum(list []int) int {
	var sum int
	for _, v := range list {
		if v%2 == 0 {
			sum += v
		}
	}
	return sum
}

// Reverse a String
func stringReverse(s string) string {
	var reversedString string
	for i := len(s) - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}
	return reversedString
}

// Palindrome Checker
func isPalindrome(s string) bool {
	var reversedString string
	for i := len(s) - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}
	return (s == reversedString)
}

// Fibonacci Sequence
func fibonacci(count int) {
	num1, num2 := 0, 1
	fmt.Println(num1)
	fmt.Println(num2)
	var num3 int
	for range count - 2 {
		num3 = num1 + num2
		fmt.Println(num3)
		num1 = num2
		num2 = num3
	}
}

// Factorial Calculation
func factorial(num int) int {
	if num <= 0 {
		return 0
	}
	factorial := 1
	for i := num; i > 0; i-- {
		factorial = factorial * i
	}
	return factorial
}

// Prime Number Checker
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
