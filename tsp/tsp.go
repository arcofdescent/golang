// Travelling salesman problem
// ---------------------------
// Given a list of cities and the distances between each pair of cities, what is
// the shortest possible route that visits each city exactly once and returns to
// the origin city?
package main

import (
	"fmt"
	"sort"
	"strings"
)

var cities = [...]string{"Bangalore", "Delhi", "Mumbai"}

// set the distances
var distance = map[string]float32{
	"del_mum": 56,
	"ban_mum": 26,
	"ban_del": 96,
}

func main() {

	fmt.Println()

	/*
		vals := [][]int{}
		res := []int{0, 1, 2}
		vals = append(vals, res)
		res = []int{0, 2, 1}
		vals = append(vals, res)
		fmt.Printf("%v\n", vals)
	*/

	vals := permutations([]int{0, 1, 2}, 3)
	fmt.Printf("%v\n", vals)

	/*
		var dst float32
		var min float32
		var minCity string

		for _, c := range cities {
			dst = calcDistance(c)
			fmt.Printf("Origin: %s, Distance: %.2f\n", c, dst)

			if min == 0 {
				min = dst
				minCity = c
			}

			if min > 0 && dst < min {
				min = dst
				minCity = c
			}
		}

		fmt.Println()
		fmt.Printf("Origin city should be %s\n", minCity)
	*/
}

// taken from Python's lib
func permutations(iterable []int, r int) [][]int {

	values := [][]int{}

	pool := iterable
	n := len(pool)

	if r > n {
		return values
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	fmt.Println(result)
	values = append(values, result)
	fmt.Printf("values: %v\n", values)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i -= 1 {
			cycles[i] -= 1
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j += 1 {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				for k := i; k < r; k += 1 {
					result[k] = pool[indices[k]]
				}

				fmt.Printf("result: %v\n", result)
				fmt.Printf("values before append: %v\n", values)
				values = append(values, result)
				fmt.Printf("values: %v\n", values)

				break
			}
		}

		if i < 0 {
			return values
		}

	}

	return values
}

func calcDistance(city string) float32 {

	var dst float32

	for _, c := range cities {
		if c == city {
			continue
		}

		dst += getDistance(city, c) * 2
	}

	return dst
}

func getDistance(city1 string, city2 string) float32 {

	// get first 3 letters
	city1 = strings.ToLower(city1[:3])
	city2 = strings.ToLower(city2[:3])

	// create a slice for sorting
	cities := []string{city1, city2}
	sort.Strings(cities)

	// key into the distance map
	dist_key := strings.Join(cities, "_")
	return distance[dist_key]
}
