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

type List[T any] struct {
	head, tail *node[T]
}

type node[T any] struct {
	next *node[T]
	val  T
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &node[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &node[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) AsSlice() []T {
	ans := make([]T, 0)
	curr := list.head
	for {
		if curr == nil {
			break
		}
		ans = append(ans, curr.val)
		curr = curr.next
	}
	return ans
}

func main() {
	m := map[int]string{1: "2", 2: "4", 4: "8"}
	k1 := MapKeys(m)
	fmt.Println(k1)
	k := MapKeys[int, string](m)
	fmt.Println(k)

	l := List[string]{}
	fmt.Println(l.AsSlice())
	l.Push("One")
	fmt.Println(l.AsSlice())
	l.Push("Two")
	fmt.Println(l.AsSlice())
	l.Push("Three")
	fmt.Println(l.AsSlice())
}
