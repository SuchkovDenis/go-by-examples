package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	fmt.Println()
	r := make([]K, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}
	return r
}

// TODO realize list

func main() {
	m := map[int]string{1: "2", 2: "4", 4: "8"}
	k1 := MapKeys(m)
	fmt.Println(k1)
	k := MapKeys[int, string](m)
	fmt.Println(k)
}
