package challenges

import "fmt"

func PanicMain() {
	defer handlePanic()
	panic("raising a panic")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from:", r)
	}
}
