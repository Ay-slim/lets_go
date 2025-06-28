package main

import "fmt"
import "errors"

func teaFactory(val int) (int, error) {
	if val == 42 {
		// It is a convention to return any possible errors alongside a regular return value
		return -1, errors.New("Cannot work with 42")
	}
	return val, nil
}

func main() {
	test_inputs := []int{3, 5, 42, 8}
	for input := range(test_inputs) {
		value, err := teaFactory(test_inputs[input])
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(value)
		}
	}
}
