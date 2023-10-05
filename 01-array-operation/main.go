package main

import (
	"fmt"
	"github.com/rahmaninsani/dipay-backend-technical-test/01-array-operation/helper"
)

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	result1 := helper.DuplicateAndShift(arr1)
	fmt.Println("Input: ", arr1)
	fmt.Println("Output: ", result1)

	arr2 := []int{1, 2, 3}
	result2 := helper.DuplicateAndShift(arr2)
	fmt.Println("\nInput: ", arr2)
	fmt.Println("Output: ", result2)
}
