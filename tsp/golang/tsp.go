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

var cities = [...]string{"Bangalore", "Delhi", "Mumbai", "Chennai"}

// set the distances
var distance = map[string]float32{
	"del_mum": 56,
	"ban_mum": 26,
	"ban_del": 96,
	"ban_che": 16,
	"che_del": 116,
	"che_mum": 36,
}

func main() {

	fmt.Println()

	vals := permutations([]int{0, 1, 2, 3})
	fmt.Printf("%v\n", vals)

	var dst float32
	/*
		var min float32
		var minCity string
	*/

	for _, route := range vals {
		fmt.Printf("Route: %s, %s, %s, %s\n",
			cities[route[0]], cities[route[1]], cities[route[2]], cities[route[3]])
		dst = 0

		for idx, _ := range route {
			if idx == len(route)-1 {
				city1 := cities[route[idx]]
				city2 := cities[route[0]]
				dst_cities := getDistance(city1, city2)
				dst += dst_cities
				fmt.Printf("city1: %s, city2: %s, distance: %.2f\n", city1, city2, dst_cities)
			} else {
				city1 := cities[route[idx]]
				city2 := cities[route[idx+1]]
				dst_cities := getDistance(city1, city2)
				dst += dst_cities
				fmt.Printf("city1: %s, city2: %s, distance: %.2f\n", city1, city2, dst_cities)
			}
		}
		/*
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
		*/
		fmt.Printf("dst: %.2f\n", dst)
	}

	/*
		fmt.Println()
		fmt.Printf("Origin city should be %s\n", minCity)
	*/
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(arr, len(arr))
	return res
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
