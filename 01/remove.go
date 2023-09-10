package main

import (
	"fmt"
	"reflect"
	"testing"
)

func is_even(number int) (r bool) {
	r = number%2 == 0
	return
}

func customFilter(list []int, fun func(number int) bool) (results []int) {
	for _, val := range list {
		if fun(val) {
			results = append(results, val)
		}
	}

	return
}

// No built-in function for Go?

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

	testNumbers := []int{1, 2, 3, 5, 6, 4, 7, 8, 9, 22, 12, 10}
	rightResponse := []int{2, 6, 4, 8, 22, 12, 10}

	if !reflect.DeepEqual(customFilter(testNumbers, is_even), rightResponse) {
		t.Errorf("customFilter(%v) = rightResponse; want true", testNumbers)
	}

	fmt.Println("Tests passed")
}

func main() {
	numbers := []int{1, 2, 3, 5, 6, 4, 7, 8, 9, 22, 12, 10}

	evenNumbersFromCustom := customFilter(numbers, is_even)

	fmt.Printf("From custom: %v\n", evenNumbersFromCustom)

	t := new(testing.T)
	ExecuteTests(t)
}
