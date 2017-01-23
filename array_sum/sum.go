package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter array elems separated by space: ")
	elem_string, _ := reader.ReadString('\n')
	elem_string = strings.TrimRight(elem_string, "\n")

	elems := strings.Split(elem_string, " ")
	fmt.Println(elems)

	// sum
	sum := 0
	for _, elem := range elems {
		val, _ := strconv.Atoi(elem)
		sum += val
	}

	fmt.Printf("sum: %d\n", sum)
}
