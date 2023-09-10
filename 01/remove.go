package main

import (
	"fmt"
	"testing"
)

func is_even(number int) (r bool) {
	r = number%2 == 0
	return
}

func ExecuteTests(t *testing.T) {
	if !is_even(2) {
		t.Errorf("is_even(2) = false; want true")
	}
	if is_even(3) {
		t.Errorf("is_even(3) = true; want false")
	}
	if !is_even(0) {
		t.Errorf("is_even(0) = false; want true")
	}

	fmt.Println("Tests passed")
}

func main() {

	numbers := [...]int{1, 2, 3, 5, 6, 4, 7, 8, 9, 22, 12, 10}

	// evenNumbersFromBuiltIn := Choose(numbers, is_even)

	fmt.Printf("From built-in: %v\n", numbers)

	t := new(testing.T)
	ExecuteTests(t)
}
