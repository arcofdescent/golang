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

	elem_string, _ := reader.ReadString('\n')
	elem_string = strings.TrimRight(elem_string, "\n")

	elems := strings.Split(elem_string, " ")

	// sum
	sum := 0
	for _, elem := range elems {
		val, _ := strconv.Atoi(elem)
		sum += val
	}

	fmt.Println(sum)
}
