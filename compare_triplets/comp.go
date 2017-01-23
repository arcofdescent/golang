package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {

	reader := bufio.NewReader(os.Stdin)

	a_str, _ := reader.ReadString('\n')
	a_str = strings.TrimRight(a_str, "\n")

	b_str, _ := reader.ReadString('\n')
	b_str = strings.TrimRight(b_str, "\n")

	a_strs := strings.Split(a_str, " ")
	b_strs := strings.Split(b_str, " ")

	a_elems := make([]int, 3)
	b_elems := make([]int, 3)

	for idx, _ := range a_strs {
		a_elems[idx], _ = strconv.Atoi(a_strs[idx])
	}

	for idx, _ := range b_strs {
		b_elems[idx], _ = strconv.Atoi(b_strs[idx])
	}

	a_score := 0
	b_score := 0

	for i := 0; i < 3; i++ {
		if a_elems[i] > b_elems[i] {
			a_score++
		} else if a_elems[i] < b_elems[i] {
			b_score++
		}
	}

	fmt.Printf("%d %d\n", a_score, b_score)
}
