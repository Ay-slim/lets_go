package main

import "fmt"

func main() {
	// arrays are fixed length data structures can be initialized with the format [length]type
	a := [5]int32{} // If initialized without values, defaults to zero
	fmt.Println(a)

	b := [5]int32{1, 2, 3, 4, 5} // You can specify on declaration liek so
	fmt.Println(b)

	c := [...]int32{1, 2, 3, 4} // Rather than specify size, you can use the ... operator to let the compiler infer size
	fmt.Println(c)

	d := [...]int32{1, 2: 800} // You can use the index to specify the value at that indexrather than typing it in directly, printed value: [1 0 800]
	fmt.Println(d)

	e := [3][3]int{} // Flex on the syntax
	for i:=0; i < 3; i++ {
		for j:=0; j < 3; j++ {
			e[i][j] = i + j
		}
	}
	fmt.Println(e)

	/*
	* Major difference between slices and arrays is that a slice does not need the size to be specified in its declaration
	*/
	s := [] int32{3, 4, 5}
	t := [] int32{}
 	fmt.Println(s, t)
	copy(t, s)
	t = append(t, 4)
	fmt.Println(s, t)
}