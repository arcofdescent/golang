package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {

	reader := bufio.NewReader(os.Stdin)

	num_elems, _ := reader.ReadString('\n')
	num_elems = strings.TrimRight(num_elems, "\n")
	//num_elems_i, _ := strconv.Atoi(num_elems)

	elem_string, _ := reader.ReadString('\n')
	elem_string = strings.TrimRight(elem_string, "\n")

	elems := strings.Split(elem_string, " ")

	arr := []int{}
	for _, elem := range elems {
		elem_i, _ := strconv.Atoi(elem)
		arr = append(arr, elem_i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d\n", arr[i])
		} else {
			fmt.Printf("%d ", arr[i])
		}
	}
}
