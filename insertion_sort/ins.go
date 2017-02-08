// implement bubble sort
package main

import "fmt"
import "math/rand"

func main() {
	var arr []int

	for i := 0; i < 30; i++ {
		arr = append(arr, rand.Intn(100))
	}

	insSort(arr)
	fmt.Printf("sorted_arr: %v\n", arr)
}

func insSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for ; j > 0 && arr[j-1] > arr[j]; j-- {
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
		}
	}
}
