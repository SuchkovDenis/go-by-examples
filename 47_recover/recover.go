package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	myPanic()
	fmt.Println("After myPanic")
}

func myPanic() {
	panic("a problem")
}
