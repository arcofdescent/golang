// implement bubble sort
package main

import "fmt"
import "math/rand"

func main() {
	var arr []int

	for i := 0; i < 30; i++ {
		arr = append(arr, rand.Intn(100))
	}

	bSort(arr)
	fmt.Printf("sorted_arr: %v\n", arr)
}

func bSort(arr []int) {
	for {
		swapped := false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				tmp := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = tmp
				swapped = true
			}
		}

		// no swaps took place
		if swapped == false {
			return
		}
	}
}
