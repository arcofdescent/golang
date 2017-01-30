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
