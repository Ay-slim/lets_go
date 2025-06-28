package main

import "fmt"
// Structs, think of these as objects or classes with properties and methods

type Node[T any] struct {
	val T
	next *Node[T]
}

type List[T any] struct {
	head, tail *Node[T]
}

func (list *List[T]) Add(val T) {
	if list.tail == nil {
		list.head = &Node[T]{val, nil}
		list.tail = list.head
	} else {
		list.tail.next = &Node[T]{val, nil}
		list.tail = list.tail.next
	}
}

func (list *List[T]) FetchAll()[]T {
	list_array := make([]T, 0)
	for list_idx := list.head; list_idx != nil; list_idx = list_idx.next {
		list_array = append(list_array, list_idx.val)
	}
	return list_array
}

func FetchKeys[K comparable, V any](arg map[K]V) []K {
	key_array := make([]K, 0)
	for keys := range(arg) {
		key_array = append(key_array, keys)
	}
	return key_array
}

func main() {
	strNodeList := List[string]{}
	strNodeList.Add("Ayo")
	strNodeList.Add("Ese")
	fmt.Println(strNodeList.FetchAll())
	test_map := map[int]string {1: "a", 2: "b"}
	fmt.Println(FetchKeys((test_map)))
}
