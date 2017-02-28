/*
Consider the number 142857. We can right-rotate this number by moving the last
digit (7) to the front of it, giving us 714285.  It can be verified that
714285=5×142857.  This demonstrates an unusual property of 142857: it is a
divisor of its right-rotation.

Find the last 5 digits of the sum of all integers n, 10 < n < 10100, that have
this property.
*/
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	var power float64
	fmt.Scan(&power)

	sum := 0
	max := math.Pow(10, power)

	for n := 11; n < int(max); n++ {
		if itPasses(n) {
			sum += n
		}
	}

	str := strconv.Itoa(sum)
	if len(str) > 5 {
		fmt.Println(str[len(str)-5:])
	} else {
		fmt.Println(str)
	}
}

func itPasses(n int) bool {

	n_rr := rightRotate(n)
	if n_rr%n == 0 {
		return true
	}

	return false
}

func rightRotate(n int) int {
	digits := strings.Split(strconv.Itoa(n), "")
	//fmt.Println(digits)

	for i := 0; i < 1; i++ {
		first := digits[0]
		digits[0] = digits[len(digits)-1]

		for j := len(digits) - 1; j > 1; j-- {
			digits[j] = digits[j-1]
		}

		digits[1] = first
	}

	ret, _ := strconv.Atoi(strings.Join(digits, ""))
	return ret
}
