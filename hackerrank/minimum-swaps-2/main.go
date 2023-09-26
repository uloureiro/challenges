package main

import (
	"fmt"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swaps int32
	for i := int32(0); i < int32(len(arr)); {
		// check if current value (arr[i]) is at the right position (i+1)
		if arr[i] != i+1 {
			valueAtIndex := arr[i]

			// swap values
			arr[i] = arr[valueAtIndex-1]
			arr[valueAtIndex-1] = valueAtIndex

			swaps++
		}

		// value at index is at the right position, move to the next index
		// otherwise, keep iterating over the same index until it gets the right value
		if arr[i] == i+1 {
			i++
		}
	}

	return swaps
}

func main() {
	fmt.Println(minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}))
	fmt.Println(minimumSwaps([]int32{4, 3, 2, 1}))
	fmt.Println(minimumSwaps([]int32{2, 3, 4, 1, 5}))
	fmt.Println(minimumSwaps([]int32{1, 3, 5, 2, 4, 6, 7}))
}
