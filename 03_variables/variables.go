package main

import (
	"fmt"
	"math"
)

const s = "constant"

func main() {
	fmt.Println(s)

	const pi = math.Pi

	fmt.Println(math.Sin(pi))
}
