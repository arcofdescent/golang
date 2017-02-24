package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	var numElems, numRotations int
	fmt.Scanln(&numElems, &numRotations)

	arr := make([]int, numElems)
	for i := range arr {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < numRotations; i++ {
		last := arr[numElems-1]
		arr[numElems-1] = arr[0]

		for j := 0; j < numElems-2; j++ {
			arr[j] = arr[j+1]
		}

		arr[numElems-2] = last
	}

	strs := make([]string, 0)
	for _, e := range arr {
		strs = append(strs, strconv.Itoa(e))
	}

	fmt.Println(strings.Join(strs, " "))
}
