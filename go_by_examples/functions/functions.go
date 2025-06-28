package main

import "fmt"

func basic_func (name string) string {
	return "Hello " + name
}

func multiple_return_func (num_1, num_2, num_3 int, word string) (int, string) {
	return num_1 + num_2 + num_3, word
}

func variadic_func (words ...string) string {
	ret_word := ""
	for _, word := range(words) {
		ret_word += word + " "
	}
	return ret_word
}

func closure_func () func () int {
	i := 0
	return func () int {
		i++
		return i
	}
}

func factorial (base int) int {
	if base == 0 {
		return 1
	}
	return base * factorial(base - 1)
} 

func main() {
	// Basic function returning a greeting
	fmt.Println(basic_func("Ayo"))

	// Function with multiple returns
	fmt.Println((multiple_return_func(3, 4, 5, "foo")))

	// Variadic function
	fmt.Println(variadic_func("foo", "bar", "buzz"))

	// Closure
	closure_1 := closure_func()
	fmt.Println(closure_1(), "closure_1 first")
	fmt.Println(closure_1(), "closure_1 second")
	fmt.Println(closure_1(), "closure_1 third")
	fmt.Println(closure_1(), "closure_1 fourth")
	closure_2 := closure_func()
	fmt.Println(closure_2(), "closure_2 first")

	// Recursion
	fmt.Println(factorial(5), "fact")

	var fib func (idx int) int
	fib = func(idx int) int {
		if idx == 0 || idx == 1 {
			return idx
		}
		return fib(idx - 2) + fib(idx - 1)
	}
	fmt.Println((fib(1)), "fibonacci 1")
	fmt.Println((fib(2)), "fibonacci 2")
	fmt.Println((fib(3)), "fibonacci 3")
}
